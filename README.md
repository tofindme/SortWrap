# SortWrap

#### 这是一个用来对slice来根据结构体不同成员来排序的小例子

```
type People struct {
	name   string
	age    int
	career string
	hobby  string
}

type ByNameSort struct {
	Peoples
	sort string //asc or desc?
}

func (by ByNameSort) Less(i, j int) bool {
	if by.sort == "asc" {
		return by.Peoples[i].name > by.Peoples[j].name
	} else if by.sort == "desc" {
		return by.Peoples[i].name < by.Peoples[j].name
	}
	return false
}

```
- 可以根据age,hobby来进行排序,同时可以指定asc还是desc
- 可以模仿上面类似代码来新增需要排序的字段


#### 同时借助git来熟悉一些常用命令
```
git reset --hard
git branch
...

```

