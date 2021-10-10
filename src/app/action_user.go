package app

import (
	"log"

	"github.com/skratchdot/open-golang/open"
	"github.com/sqweek/dialog"
	"sdmm/app/ui/cpwsarea/workspace/wsmap"
	"sdmm/dm/dmmap"
	"sdmm/dm/dmmap/dmmdata"
)

/*
	File similar to action.go, but contains methods triggered by user. (ex. when button clicked)
	Such methods have a "Do" prefix, and they are logged excessively.
*/

// AppDoOpenEnvironment opens environment, which user need to select in file dialog.
func (a *app) AppDoOpenEnvironment() {
	log.Println("[app] selecting environment to open...")
	if file, err := dialog.
		File().
		Title("Open Environment").
		Filter("*.dme", "dme").
		Load(); err == nil {
		log.Println("[app] environment to open selected:", file)
		a.openEnvironment(file)
	}
}

// AppDoOpenEnvironmentByPath opens environment by provided path.
func (a *app) AppDoOpenEnvironmentByPath(path string) {
	log.Println("[app] open environment by path:", path)
	a.openEnvironment(path)
}

// AppDoClearRecentEnvironments clears recently opened environments.
func (a *app) AppDoClearRecentEnvironments() {
	log.Println("[app] clear recent environments")
	a.configData.ClearRecentEnvironments()
	a.configData.Save()
}

// AppDoSelectMapFile opens dialog window to select a map file.
func (a *app) AppDoSelectMapFile() (string, error) {
	log.Println("[app] selecting map file...")
	return dialog.
		File().
		Title("Open Map").
		Filter("*.dmm", "dmm").
		SetStartDir(a.loadedEnvironment.RootDir).
		Load()
}

// AppDoOpenMap opens map, which user need to select in file dialog.
func (a *app) AppDoOpenMap() {
	log.Println("[app] selecting map to open...")
	if file, err := a.AppDoSelectMapFile(); err == nil {
		log.Println("[app] map to open selected:", file)
		a.openMap(file)
	}
}

// AppDoOpenMapByPath opens map by provided path.
func (a *app) AppDoOpenMapByPath(path string) {
	log.Println("[app] open map by path:", path)
	a.openMap(path)
}

// AppDoOpenMapByPathV same as AppDoOpenMapByPath by map will be opened inside the concrete workspace with the provided index.
func (a *app) AppDoOpenMapByPathV(path string, workspaceIdx int) {
	log.Printf("[app] open map with workspace index [%d] by path: [%s]", workspaceIdx, path)
	a.openMapV(path, workspaceIdx)
}

// AppDoClearRecentMaps clears recently opened maps.
func (a *app) AppDoClearRecentMaps() {
	log.Println("[app] clear recent maps")
	a.configData.ClearRecentMaps(a.loadedEnvironment.RootFile)
	a.configData.Save()
}

// AppDoSave saves current active map.
func (a *app) AppDoSave() {
	log.Println("[app] do save")
	if activeWs := a.layout.WsArea.ActiveWorkspace(); activeWs != nil {
		if activeWs, ok := activeWs.(*wsmap.WsMap); ok {
			activeWs.Save()
		}
	}
}

// AppDoOpenPreferences opens preferences tab.
func (a *app) AppDoOpenPreferences() {
	log.Println("[app] open preferences")
	a.layout.OpenPreferences(a.makePreferences())
}

// AppDoSelectInstance globally selects provided instance in the app.
func (a *app) AppDoSelectInstance(instance *dmmdata.Instance) {
	log.Printf("[app] select instance: path=[%s], id=[%d]", instance.Path(), instance.Id())
	a.layout.Environment.SelectPath(instance.Path())
	a.layout.Instances.Select(instance)
}

// AppDoSelectInstanceByPath globally selects an instance with provided type path.
func (a *app) AppDoSelectInstanceByPath(path string) {
	log.Println("[app] select instance by path:", path)
	a.AppDoSelectInstance(dmmap.InstanceCache.Get(path, a.AppInitialInstanceVariables(path)))
}

// AppDoExit exits the app.
func (a *app) AppDoExit() {
	log.Println("[app] exit")
	a.tmpShouldClose = true
}

// AppDoUndo does undo of the latest command.
func (a *app) AppDoUndo() {
	log.Println("[app] undo")
	a.commandStorage.Undo()
}

// AppDoRedo does redo of the previous command.
func (a *app) AppDoRedo() {
	log.Println("[app] redo")
	a.commandStorage.Redo()
}

// AppDoResetWindows resets application windows to their initial positions.
func (a *app) AppDoResetWindows() {
	log.Println("[app] reset windows")
	a.resetWindows()
}

// AppDoOpenLogs open the logs folder.
func (a *app) AppDoOpenLogs() {
	log.Println("[app] open logs dir:", a.logDir)
	if err := open.Run(a.logDir); err != nil {
		log.Println("[app] unable to open log dir:", err)
	}
}
