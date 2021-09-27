// Package concmap is modelled after a sample given on SO.
// The DataStore type was not specified so a sample is included.
// Reference: https://stackoverflow.com/a/48956497
package concmap

// DataStore is a sample data storage implementation.
// Update this type to reflect specific needs.
type DataStore struct{}

func (s *DataStore) GetId() string {
	return "fake"
}

type DataManager struct {
	/** This contain connection to know dataStore **/
	m_dataStores map[string]DataStore

	/** That channel is use to access the dataStores map **/
	m_dataStoreChan chan map[string]interface{}
}

func newDataManager() *DataManager {
	dataManager := new(DataManager)
	dataManager.m_dataStores = make(map[string]DataStore)
	dataManager.m_dataStoreChan = make(chan map[string]interface{}, 0)
	// Concurrency...
	go func() {
		for {
			select {
			case op := <-dataManager.m_dataStoreChan:
				if op["op"] == "getDataStore" {
					storeId := op["storeId"].(string)
					op["store"].(chan DataStore) <- dataManager.m_dataStores[storeId]
				} else if op["op"] == "getDataStores" {
					stores := make([]DataStore, 0)
					for _, store := range dataManager.m_dataStores {
						stores = append(stores, store)
					}
					op["stores"].(chan []DataStore) <- stores
				} else if op["op"] == "setDataStore" {
					store := op["store"].(DataStore)
					dataManager.m_dataStores[store.GetId()] = store
				} else if op["op"] == "removeDataStore" {
					storeId := op["storeId"].(string)
					delete(dataManager.m_dataStores, storeId)
				}
			}
		}
	}()

	return dataManager
}

/**
 * Access Map functions...
 */
func (this *DataManager) getDataStore(id string) DataStore {
	arguments := make(map[string]interface{})
	arguments["op"] = "getDataStore"
	arguments["storeId"] = id
	result := make(chan DataStore)
	arguments["store"] = result
	this.m_dataStoreChan <- arguments
	return <-result
}

func (this *DataManager) getDataStores() []DataStore {
	arguments := make(map[string]interface{})
	arguments["op"] = "getDataStores"
	result := make(chan []DataStore)
	arguments["stores"] = result
	this.m_dataStoreChan <- arguments
	return <-result
}

func (this *DataManager) setDataStore(store DataStore) {
	arguments := make(map[string]interface{})
	arguments["op"] = "setDataStore"
	arguments["store"] = store
	this.m_dataStoreChan <- arguments
}

func (this *DataManager) removeDataStore(id string) {
	arguments := make(map[string]interface{})
	arguments["storeId"] = id
	arguments["op"] = "removeDataStore"
	this.m_dataStoreChan <- arguments
}
