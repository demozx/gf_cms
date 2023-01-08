package packed

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sChannel struct{}

var (
	insChannel = sChannel{}
)

func init() {
	service.RegisterChannel(New())
}

func New() *sChannel {
	return &sChannel{}
}

func Channel() *sChannel {
	return &insChannel
}

// BackendApiIndex 获取后台栏目分类接口数据
func (s *sChannel) BackendApiIndex(ctx context.Context) (out []*model.ChannelBackendApiListItem, err error) {
	var allChannels []*entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).OrderAsc(dao.CmsChannel.Columns().Sort).OrderAsc(dao.CmsChannel.Columns().Id).Scan(&allChannels)
	if err != nil {
		return nil, err
	}
	var channelBackendApiList []*model.ChannelBackendApiListItem
	for key, channel := range allChannels {
		if channel.Thumb != "" {
			channel.Name = channel.Name + "<span style='color:red;font-size: 12px;margin: 0 2px;'>图</span><span id='id_" + gconv.String(channel.Id) + "' class='cate_id'>&nbsp;ID:" + gconv.String(channel.Id) + "</span>"

		} else {
			channel.Name = channel.Name + "<span id='id_" + gconv.String(channel.Id) + "' class='cate_id'>&nbsp;ID:" + gconv.String(channel.Id) + "</span>"
		}
		allChannels[key] = channel
	}
	err = gconv.Scan(allChannels, &channelBackendApiList)
	if err != nil {
		return nil, err
	}
	channelBackendApiList = Channel().channelBackendApiListRecursion(channelBackendApiList, 0)
	return channelBackendApiList, nil
}

func (s *sChannel) channelBackendApiListRecursion(list []*model.ChannelBackendApiListItem, pid int) (out []*model.ChannelBackendApiListItem) {
	res := make([]*model.ChannelBackendApiListItem, 0, len(list))
	for _, item := range list {
		if item.Pid == pid {
			item.Children = Channel().channelBackendApiListRecursion(list, item.Id)
			if item.Children == nil {
				item.Children = make([]*model.ChannelBackendApiListItem, 0, len(list))
			}
			item.ModelDesc = service.Channel().BackendModelDesc(item.Model)
			res = append(res, item)
		}
	}
	return res
}

// BackendChannelTree 获取栏目分类树
func (s *sChannel) BackendChannelTree(ctx context.Context, selectedId int) (out []*model.ChannelBackendApiListItem, err error) {
	var allChannels []*entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).OrderAsc(dao.CmsChannel.Columns().Sort).OrderAsc(dao.CmsChannel.Columns().Id).Scan(&allChannels)
	if err != nil {
		return nil, err
	}
	var channelBackendApiList []*model.ChannelBackendApiListItem
	err = gconv.Scan(allChannels, &channelBackendApiList)
	if err != nil {
		return nil, err
	}
	channelBackendApiList = Channel().channelBackendApiListRecursion(channelBackendApiList, 0)
	//g.Dump(channelBackendApiList)
	channelBackendApiList = Channel().backendTree(channelBackendApiList, selectedId)

	return channelBackendApiList, err
}

func (s *sChannel) BackendChannelModelTree(ctx context.Context, modelType string, channelId int) (out []*model.ChannelBackendApiListItem, err error) {
	tree, err := service.Channel().BackendChannelTree(ctx, channelId)
	if err != nil {
		return nil, err
	}
	out = make([]*model.ChannelBackendApiListItem, 0, len(tree))
	for _, item := range tree {
		if item.Model == modelType {
			out = append(out, item)
		}
	}
	return
}

// 递归生成栏目分类树
func (s *sChannel) backendTree(list []*model.ChannelBackendApiListItem, selectedPid int) (out []*model.ChannelBackendApiListItem) {
	var hasChildren = false
	newList := make([]*model.ChannelBackendApiListItem, 0, len(list))
	for _, item := range list {
		newItem := new(model.ChannelBackendApiListItem)
		newItem.Id = item.Id
		newItem.Pid = item.Pid
		newItem.Level = item.Level
		newItem.Status = item.Status
		newItem.Name = item.Name
		newItem.Model = item.Model
		newItem.Children = nil
		newList = append(newList, newItem)
		if len(item.Children) > 0 {
			hasChildren = true
			for _, childrenItem := range item.Children {
				newList = append(newList, childrenItem)
			}
		}
	}
	if hasChildren == true {
		return Channel().backendTree(newList, selectedPid)
	}
	for key, item := range newList {
		var emsp = ""
		for i := 0; i < item.Level; i++ {
			emsp += "└"
		}
		newList[key].Name = "" + emsp + "&nbsp;" + item.Name
		if item.Id == selectedPid {
			newList[key].Selected = 1
		}
	}
	return newList
}

// BackendApiStatus 状态禁用启用
func (s *sChannel) BackendApiStatus(ctx context.Context, in *backendApi.ChannelStatusApiReq) (out *backendApi.ChannelStatusApiRes, err error) {
	var first *entity.CmsChannel
	var m = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, in.Id)
	err = m.Scan(&first)
	if err != nil {
		return nil, err
	}
	if first == nil {
		return nil, gerror.New("栏目不存在")
	}
	status := 1
	if first.Status == 1 {
		status = 0
	}
	_, err = m.Data(g.Map{
		dao.CmsChannel.Columns().Status: status,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sChannel) BackendApiDelete(ctx context.Context, in *backendApi.ChannelDeleteApiReq) (out *backendApi.ChannelDeleteApiRes, err error) {
	var first *entity.CmsChannel
	var m = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, in.Id)
	err = m.Scan(&first)
	if err != nil {
		return nil, err
	}
	if first == nil {
		return nil, gerror.New("栏目不存在")
	}
	var children *entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Pid, in.Id).Scan(&children)
	if err != nil {
		return nil, err
	}
	if children != nil {
		return nil, gerror.New("当前栏目下有子栏目，不允许删除")
	}
	// todo 栏目下有内容也不能删除
	_, err = m.Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sChannel) BackendApiAdd(ctx context.Context, in *backendApi.ChannelAddApiReq) (out *backendApi.ChannelAddApiRes, err error) {
	if len(in.ListTemplate) > 0 && !gstr.HasPrefix(in.ListTemplate, "{module}") {
		return nil, gerror.New("频道模板需以'{module}'开头，以便自动获取模块")
	}
	if len(in.DetailTemplate) > 0 && !gstr.HasPrefix(in.DetailTemplate, "{module}") {
		return nil, gerror.New("详情模板需以'{module}'开头，以便自动获取模块")
	}
	var entityData *entity.CmsChannel
	err = gconv.Scan(in, &entityData)
	if err != nil {
		return nil, err
	}
	var parent *entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, in.Pid).Scan(&parent)
	if err != nil {
		return nil, err
	}
	entityData.Level = 0
	if parent != nil {
		entityData.Level = parent.Level + 1
	}
	id, err := dao.CmsChannel.Ctx(ctx).Data(entityData).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	_, err = service.Channel().UpdateRelation(ctx, gconv.Int(id))
	if err != nil {
		return nil, err
	}
	return
}

func (s *sChannel) BackendApiEdit(ctx context.Context, in *backendApi.ChannelEditApiReq) (out *backendApi.ChannelEditApiRes, err error) {
	var oldLevel = 0
	var newLevel = 0
	var newPid = 0
	if in.Pid > 0 {
		parent, err := Channel().GetOneById(ctx, in.Pid)
		if err != nil {
			return nil, err
		}
		oldLevel = parent.Level
		newLevel = parent.Level + 1
		newPid = gvar.New(parent.Id).Int()
	}
	channel, err := Channel().GetOneById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if channel.Pid != newPid && oldLevel >= channel.Level {
		return nil, gerror.New("不能选择同级别或级别在自己下边的分类")
	}
	if in.Pid == in.Id {
		return nil, gerror.New("自己不能是自己的父级分类")
	}
	if len(in.ListTemplate) > 0 && !gstr.HasPrefix(in.ListTemplate, "{module}") {
		return nil, gerror.New("频道模板需以'{module}'开头，以便自动获取模块")
	}
	if len(in.DetailTemplate) > 0 && !gstr.HasPrefix(in.DetailTemplate, "{module}") {
		return nil, gerror.New("详情模板需以'{module}'开头，以便自动获取模块")
	}
	var data *entity.CmsChannel
	err = gconv.Scan(in, &data)
	if err != nil {
		return nil, err
	}
	data.Level = newLevel
	data.Pid = newPid
	_, err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, in.Id).Update(data)
	if err != nil {
		return nil, err
	}
	_, err = service.Channel().UpdateRelation(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sChannel) GetOneById(ctx context.Context, id int) (out *entity.CmsChannel, err error) {
	var channel *entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, id).Scan(&channel)
	if err != nil {
		return nil, err
	}
	if channel == nil {
		return nil, gerror.New("栏目分类不存在")
	}
	return channel, nil
}

// BackendModelMap 后台数据模型映射
func (s *sChannel) BackendModelMap() map[string]string {
	return map[string]string{
		consts.ChannelModelArticle:    consts.ChannelModelArticleDesc,
		consts.ChannelModelImage:      consts.ChannelModelImageDesc,
		consts.ChannelModelSinglePage: consts.ChannelModelSinglePageDesc,
	}
}

// BackendModelCanAddMap 后台允许填充列表内容的数据模型映射
func (s *sChannel) BackendModelCanAddMap() map[string]string {
	return map[string]string{
		consts.ChannelModelArticle: consts.ChannelModelArticleDesc,
		consts.ChannelModelImage:   consts.ChannelModelImageDesc,
	}
}

func (s *sChannel) BackendModelDesc(model string) string {
	modelMap := Channel().BackendModelMap()
	out, isOk := modelMap[model]
	if isOk == false {
		return ""
	}
	return out
}

// UpdateRelation 更新关联关系字段
func (s *sChannel) UpdateRelation(ctx context.Context, originChannelId int) (out interface{}, err error) {
	var chUpdateRelation = make(chan int, 2)
	go func() {
		_, err := Channel().updateTid(ctx, originChannelId, 0)
		if err != nil {
			return
		}
		chUpdateRelation <- 1
	}()
	<-chUpdateRelation
	channelTid := 0
	value, err := dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, originChannelId).Value(dao.CmsChannel.Columns().Tid)
	if err != nil {
		return nil, err
	}
	if !value.IsEmpty() {
		channelTid = value.Int()
	}
	go func() {
		_, err := Channel().updateChildren(ctx, channelTid, []int{}, []int{}, 0)
		if err != nil {
			return
		}
		chUpdateRelation <- 1
	}()
	<-chUpdateRelation
	defer close(chUpdateRelation)
	return
}

// UpdateTid 更新tid
func (s *sChannel) updateTid(ctx context.Context, originChannelId int, pid int) (out interface{}, err error) {
	tid := originChannelId
	if pid > 0 {
		tid = pid
	}
	var channelInfo *entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, tid).Scan(&channelInfo)
	if err != nil {
		return false, err
	}
	if channelInfo == nil {
		return nil, gerror.New("栏目不存在")
	}
	if channelInfo.Pid == 0 {
		// 自己就是顶级了，所以顶级是自己
		_, err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, originChannelId).Data(g.Map{
			dao.CmsChannel.Columns().Tid: tid,
		}).Update()
		if err != nil {
			return nil, err
		}
	} else {
		_, err = Channel().updateTid(ctx, originChannelId, channelInfo.Pid)
		if err != nil {
			return nil, err
		}
	}
	return
}

func (s *sChannel) updateChildren(ctx context.Context, channelTid int, lastBatchIdsArr []int, allIdsArr []int, level int) (out interface{}, err error) {
	//g.Dump("idsArr", lastBatchIdsArr)
	//g.Dump("allIdsArr", allIdsArr)
	//g.Dump("level", level)
	var array []*gvar.Var
	if level == 0 {
		array, err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Pid, channelTid).Array(dao.CmsChannel.Columns().Id)
	} else {
		array, err = dao.CmsChannel.Ctx(ctx).WhereIn(dao.CmsChannel.Columns().Pid, lastBatchIdsArr).Array(dao.CmsChannel.Columns().Id)
	}
	if err != nil {
		return nil, err
	}
	//g.Dump("array", array)
	//time.Sleep(time.Second * 2)
	var arrayInt []int
	for _, id := range array {
		arrayInt = append(arrayInt, gconv.Int(id))
	}
	//g.Dump("arrayInt", arrayInt)
	//time.Sleep(time.Second * 2)
	if level == 0 {
		// 第一次
		//g.Dump("第一次递归")
		_, err = Channel().updateChildren(ctx, channelTid, arrayInt, arrayInt, level+1)
		if err != nil {
			return nil, err
		}
	}
	if level != 0 && len(arrayInt) == 0 {
		// 最后一次递归
		//g.Dump("最后一次递归")
		//g.Dump("allIdsArr", allIdsArr)
		var arrayStr []string
		for _, id := range allIdsArr {
			arrayStr = append(arrayStr, gconv.String(id))
		}
		//g.Dump("arrayStr", arrayStr)
		_, err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, channelTid).Data(g.Map{
			dao.CmsChannel.Columns().ChildrenIds: gstr.Implode(",", arrayStr),
		}).Update()
		if err != nil {
			return nil, err
		}
		return
	} else {
		for _, id := range arrayInt {
			allIdsArr = append(allIdsArr, id)
		}
		//g.Dump("allIdsArr", allIdsArr)
		_, err = Channel().updateChildren(ctx, channelTid, arrayInt, allIdsArr, level+1)
		if err != nil {
			return nil, err
		}
	}
	return
}

// GetChildIds 获取当前栏目的所有子栏目id
// belongChannelId 当前栏目
// andMe 是否包含自身id
func (s *sChannel) GetChildIds(ctx context.Context, belongChannelId int, andMe bool) (arrAllIds []int, err error) {
	// 获取当前栏目的所有子栏目
	strChildrenIds, err := dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, belongChannelId).Value(dao.CmsChannel.Columns().ChildrenIds)
	if err != nil {
		return nil, err
	}
	// 将所有子栏目转成数组
	arrChildrenIds := gstr.SplitAndTrim(strChildrenIds.String(), ",")
	if belongChannelId > 0 && andMe {
		// 将当前指定的最完成栏目id存进数组
		arrAllIds = append(arrAllIds, belongChannelId)
	}
	if len(arrChildrenIds) > 0 {
		for _, id := range arrChildrenIds {
			// 将子栏目id们存进数组
			arrAllIds = append(arrAllIds, gconv.Int(id))
		}
	}
	return
}
