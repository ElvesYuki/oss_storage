package idgenerator

import (
	"go.uber.org/zap"
	"math"
	"oss_storage/common/httperror"
)

const (
	MODULE_DEFAULT   = "default"
	MODULE_TEMP      = "temp"
	MODULE_TEST      = "test"
	MODULE_OSS_EVENT = "oss_event"
)

var moduleMap map[string]*idModule

var updateModuleMapChan chan *idModule

type idModule struct {
	id         int64
	module     string
	step       int64
	bufferSize int64
	counter    int64
	icChan     chan int64
}

func Init() (err error) {

	sysIdCountArray, err := ListSysIdCount()
	if err != nil {
		//fmt.Println("获取module数组出错")
		return err
	}

	updateModuleMapChan = make(chan *idModule, len(sysIdCountArray))
	moduleMap = make(map[string]*idModule, len(sysIdCountArray))

	for _, sysIdCount := range sysIdCountArray {
		// 构造 moduleMap
		sysIdModule := new(idModule)
		sysIdModule.id = sysIdCount.Id.Int64
		sysIdModule.module = sysIdCount.Module.String
		sysIdModule.step = sysIdCount.Step.Int64
		sysIdModule.bufferSize = int64(math.Ceil(float64(sysIdCount.Step.Int64) * 1.2))
		sysIdModule.counter = sysIdCount.Counter.Int64
		sysIdModule.icChan = make(chan int64, sysIdModule.bufferSize)

		moduleMap[sysIdModule.module] = sysIdModule

		// 用channel更新
		updateModuleMapChan <- sysIdModule
		go updateSysIdModule()
	}
	return nil
}

// addSysIdModule 新增本地的ModuleMap
func addSysIdModule(module string) (*idModule, error) {

	var sysIdModule *idModule
	sysIdCountArray, err := ListSysIdCount()
	if err != nil {
		err = new(httperror.XmoError).WithBiz(httperror.BIZ_SQL_NOT_EXIST_ERROR)
		return nil, err
	}

	for _, sysIdCount := range sysIdCountArray {
		if module == sysIdCount.Module.String {
			sysIdModule = new(idModule)
			sysIdModule.id = sysIdCount.Id.Int64
			sysIdModule.module = sysIdCount.Module.String
			sysIdModule.step = sysIdCount.Step.Int64
			sysIdModule.bufferSize = int64(math.Ceil(float64(sysIdCount.Step.Int64) * 1.2))
			sysIdModule.counter = sysIdCount.Counter.Int64
			sysIdModule.icChan = make(chan int64, sysIdModule.bufferSize)
		}
	}

	if sysIdModule == nil {
		// 抛出异常
		err = new(httperror.XmoError).WithBiz(httperror.BIZ_SQL_NOT_EXIST_ERROR)
		return nil, err
	}

	// 用channel更新
	go updateSysIdModule()
	updateModuleMapChan <- sysIdModule

	return sysIdModule, err
}

// updateSysIdModule 更新Map, 刷新idChan
func updateSysIdModule() {
	// TODO 分布式锁开始

	sysIdModule := <-updateModuleMapChan

	//fmt.Println("更新chan", sysIdModule)
	moduleMap[sysIdModule.module] = sysIdModule

	// 访问数据库开始
	// 获取最新的counter
	sysIdCount, err := GetSysIdCountById(sysIdModule.id)
	if err != nil {
		zap.L().Error(httperror.BIZ_SQL_NOT_EXIST_ERROR.Msg, zap.Error(err))
		panic(err.Error())
	}
	// 更新本地counter
	sysIdModule.step = sysIdCount.Step.Int64
	sysIdModule.bufferSize = int64(math.Ceil(float64(sysIdCount.Step.Int64) * 1.2))
	sysIdModule.counter = sysIdCount.Counter.Int64

	// 用channel更新
	go updateModuleMap()
	updateModuleMapChan <- sysIdModule

	// 更新数据库counter
	if err := UpdateCounterSysIdCountById(sysIdModule.id, sysIdModule.counter+sysIdModule.step); err != nil {
		zap.L().Error(httperror.BIZ_SQL_UPDATE_ERROR.Msg, zap.Error(err))
		panic(err.Error())
	}
	// 访问数据库结束

	// 生成 连续Id
	if int64(cap(sysIdModule.icChan)) != sysIdModule.bufferSize {
		close(sysIdModule.icChan)
		//fmt.Println("不相等")
		sysIdModule.icChan = make(chan int64, sysIdModule.bufferSize)
	}

	for i := int64(0); i < sysIdModule.step; i++ {
		sysIdModule.icChan <- sysIdModule.counter + i
	}
	// TODO 分布式锁结束
}

// updateModuleMap 更新本地的ModuleMap
func updateModuleMap() {
	// TODO 写锁
	updateIdModule := <-updateModuleMapChan
	moduleMap[updateIdModule.module] = updateIdModule
}

// checkIdChan 检查是否需要更新 idChan
func checkIdChan(sysIdModule *idModule) {
	if int64(len(sysIdModule.icChan)) == (sysIdModule.bufferSize - sysIdModule.step) {
		// 用channel更新
		updateModuleMapChan <- sysIdModule
		go updateSysIdModule()
	}
}

func GetId() (int64, error) {
	return GetIdByModule(MODULE_DEFAULT)
}

func GetIdByModule(module string) (id int64, err error) {

	var sysIdModule *idModule
	var hasModule bool

	// TODO  获取写锁
	sysIdModule, hasModule = moduleMap[module]
	if !hasModule {
		sysIdModule, err = addSysIdModule(module)
		if err != nil {
			return 0, err
		}
	}
	go checkIdChan(sysIdModule)

	id = <-sysIdModule.icChan

	if id == 0 {
		id, err = GetIdByModule(module)
	}

	//fmt.Println("生成的Id===>",strconv.FormatInt(id, 10))

	return id, err
}
