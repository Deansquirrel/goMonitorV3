package taskService

type HealthTask struct {
}

//
//func (ht *HealthTask) StartTask() error {
//	healthTaskConfigRepository := taskConfigRepository.HealthTaskConfig{}
//	//获取配置列表
//	list, err := healthTaskConfigRepository.GetHealthConfigList()
//	if err != nil {
//		return err
//	}
//	//清空Config列表
//	healthConfigList = make(map[string]*taskConfigRepository.HealthTaskConfigData)
//	//清空Task列表
//	if len(healthTaskList) > 0 {
//		keyList := make([]string, 0)
//		for key := range healthTaskList {
//			keyList = append(keyList, key)
//		}
//		for _, key := range keyList {
//			healthTaskList[key].Stop()
//			delete(healthConfigList, key)
//		}
//	}
//	//缓存配置列表、任务列表
//	errMsg := ""
//	errMsgFormat := "添加Health任务[%s]报错：%s；"
//	for _, config := range list {
//		err = ht.addTask(config)
//		if err != nil {
//			errMsg = errMsg + fmt.Sprintf(errMsgFormat, config.FId, err.Error())
//		}
//	}
//	if errMsg != "" {
//		log.Error(errMsg)
//		return errors.New(errMsg)
//	}
//	ht.startRegularRefresh()
//	return nil
//}
//
//func (ht *HealthTask) StartJob(id string) error {
//	c, ok := healthTaskList[id]
//	if !ok {
//		return errors.New("task is not exist")
//	}
//	c.Start()
//	ht.setTaskRunningState(id, true)
//	return nil
//}
//
//func (ht *HealthTask) StopJob(id string) error {
//	c, ok := healthTaskList[id]
//	if !ok {
//		return errors.New("task is not exist")
//	}
//	c.Stop()
//	ht.setTaskRunningState(id, false)
//	return nil
//}
//
//func (ht *HealthTask) startRegularRefresh() {
//	c := cron.New()
//	var err error
//	err = c.AddFunc("0 0/1 * * * ?", ht.refreshConfig)
//	if err != nil {
//		log.Error("添加Health配置刷新任务时遇到错误：" + err.Error())
//	} else {
//		log.Info("添加Health配置刷新任务完成")
//	}
//	c.Start()
//}
//
////刷新Health任务配置
//func (ht *HealthTask) refreshConfig() {
//	err := ht.RefreshConfig()
//	if err != nil {
//		log.Error("刷新Health配置时遇到错误：" + err.Error())
//	}
//}
//
//func (ht *HealthTask) RefreshConfig() error {
//	healthTaskConfigRepository := taskConfigRepository.HealthTaskConfig{}
//	//获取配置列表
//	list, err := healthTaskConfigRepository.GetHealthConfigList()
//	if err != nil {
//		return err
//	}
//	listId := make([]string, 0)
//	mapId := make(map[string]*taskConfigRepository.HealthTaskConfigData, 0)
//	for _, config := range list {
//		listId = append(listId, config.FId)
//		mapId[config.FId] = config
//	}
//	configId := make([]string, 0)
//	for key := range healthConfigList {
//		configId = append(configId, key)
//	}
//
//	addList, delList, checkList := goToolCommon.CheckDiff(listId, configId)
//
//	errMsg := ""
//	errMsgFormat := "刷新Health任务[%s]报错：%s；"
//
//	for _, id := range addList {
//		err = ht.addTask(mapId[id])
//		if err != nil {
//			errMsg = errMsg + fmt.Sprintf(errMsgFormat, id, err.Error())
//		}
//	}
//	for _, id := range delList {
//		ht.removeTask(id)
//	}
//	for _, id := range checkList {
//		err = ht.checkTask(mapId[id])
//		if err != nil {
//			errMsg = errMsg + fmt.Sprintf(errMsgFormat, id, err.Error())
//		}
//	}
//	if errMsg != "" {
//		return errors.New(errMsg)
//	}
//	return nil
//}
//
//func (ht *HealthTask) addTask(config *taskConfigRepository.HealthTaskConfigData) error {
//	//------------------------------------------------------------------------------------------------------------------
//	configStr, err := goToolCommon.GetJsonStr(config)
//	if err != nil {
//		log.Warn(fmt.Sprintf("Add Health Task，转换配置内容时遇到错误:%s，configID：%s", configStr, config.FId))
//	} else {
//		log.Warn(fmt.Sprintf("Add Health Task:%s", configStr))
//	}
//	//------------------------------------------------------------------------------------------------------------------
//	healthConfigList[config.FId] = config
//	w := worker.NewHealthWorker(config)
//	c := cron.New()
//	err = c.AddJob(config.FCron, w)
//	if err != nil {
//		log.Error(err.Error())
//		ht.setTaskRunningState(config.FId, false)
//	} else {
//		c.Start()
//		ht.setTaskRunningState(config.FId, true)
//	}
//	healthTaskList[config.FId] = c
//	return err
//}
//
//func (ht *HealthTask) checkTask(config *taskConfigRepository.HealthTaskConfigData) error {
//	exConfig, ok := healthConfigList[config.FId]
//	if !ok {
//		return ht.addTask(config)
//	}
//	if exConfig.IsEqual(config) {
//		return nil
//	}
//	ht.removeTask(config.FId)
//	return ht.addTask(config)
//}
//
//func (ht *HealthTask) removeTask(id string) {
//	ht.clearConfigList(id)
//	ht.clearTaskList(id)
//	ht.delTaskRunningState(id)
//}
//
//func (ht *HealthTask) clearConfigList(id string) {
//	config, ok := healthConfigList[id]
//	if !ok {
//		log.Warn(fmt.Sprintf("remove task :config is not exist,taskId[%s]", id))
//		return
//	}
//	configStr, err := goToolCommon.GetJsonStr(config)
//	if err != nil {
//		log.Warn(fmt.Sprintf("Del Health Task，转换配置内容时遇到错误:%s，configID：%s", configStr, config.FId))
//	} else {
//		log.Warn(fmt.Sprintf("Del Health Task:%s", configStr))
//	}
//	delete(healthConfigList, id)
//}
//
//func (ht *HealthTask) clearTaskList(id string) {
//	c, ok := healthTaskList[id]
//	if !ok {
//		log.Warn(fmt.Sprintf("remove task :task is not exist,taskId[%s]", id))
//		return
//	}
//	c.Stop()
//	delete(healthTaskList, id)
//}
//
//func (ht *HealthTask) setTaskRunningState(id string, s bool) {
//	healthTaskState[id] = s
//}
//
//func (ht *HealthTask) delTaskRunningState(id string) {
//	delete(healthTaskState, id)
//}
