# sensitive-go

Fork from [importcjj/sensitive](https://github.com/importcjj/sensitive).
敏感词查找，验证，过滤以及替换。内置特色10w+敏感词。[![CI-Test](https://github.com/Tohrusky/sensitive-go/actions/workflows/CI-Test.yml/badge.svg)](https://github.com/Tohrusky/sensitive-go/actions/workflows/CI-Test.yml)

## 安装

```bash
go get github.com/Tohrusky/sensitive-go
```

## 使用

```go
package sensitive

import (
	"github.com/Tohrusky/sensitive-go/sdict"
	"github.com/Tohrusky/sensitive-go/sensitive"
	"sync"
)

var (
	filter     *sensitive.Filter
	filterBoss *sensitive.Filter
	once       sync.Once
)

// Init 初始化敏感词库，单例模式
func Init() {
	once.Do(func() {
		filter = sensitive.NewWithDefaultSDict()
		filterBoss = sensitive.NewWithBossSDict()
	})
}

// ValidateBoss 重点敏感词校验，校验一个句子是否包含“重点”敏感词，合法返回true，有敏感内容返回false
func ValidateBoss(s string) bool {
	b, _ := filterBoss.Validate(s)
	return b
}

// Replace 和谐敏感词，将敏感词替换为*号
func Replace(s string) string {
	return filter.Replace(s, '*')
}
```

#### NewWithDefaultSDict & NewWithBossSDict

初始化一个携带默认敏感词的过滤器 (github.com/Tohrusky/sensitive-go/sdict)。

```go
filter := sensitive.NewWithDefaultSDict()
// filter := sensitive.NewWithBossSDict()
```

#### New

初始化一个敏感词过滤器，不携带敏感词，需要手动添加。

```go
filter := sensitive.New()
filter.AddWord(sdict.DefaultSDict...)
filter.AddWord(sdict.BossSDict...)
```

#### AddWord

添加敏感词。

```go
filter.AddWord("垃圾")
```

#### DelWord

删除敏感词。

```go
filter.DelWord("垃圾")
```

#### Filter

过滤敏感词，直接移除词语。

```go
filter.Filter("这篇文章真的好垃圾啊")
// output => 这篇文章真的好啊
```

#### Replace

和谐敏感词，把词语中的字符替换成指定的字符。

```go
filter.Replace("这篇文章真的好垃圾", '*')
// output => 这篇文章真的好**
```

#### FindIn

查找并返回第一个敏感词，如果没有则返回`false`。

```go
filter.FindIn("这篇文章真的好垃圾")
// output => true, 垃圾
```

#### FindAll

查找内容中的全部敏感词，以数组返回。

```go
filter.FindAll("这篇文章真的好垃圾")
// output => [垃圾]
```

#### Validate

验证内容是否ok，如果含有敏感词，则返回`false`和第一个敏感词。

```go
filter.Validate("这篇文章真的好垃圾")
// output => false, 垃圾
```

#### UpdateNoisePattern

设置噪音模式，排除噪音字符。

```go
// failed
filter.FindIn("这篇文章真的好垃x圾") // false
filter.UpdateNoisePattern(`x`)
// success
filter.FindIn("这篇文章真的好垃x圾") // true, 垃圾
filter.Validate("这篇文章真的好垃x圾") // False, 垃圾
```

#### LoadWordDict

加载本地敏感词字典。

```go
filter.LoadWordDict("./dict.txt")
```

#### LoadNetWordDict

加载网络词库。

```go
filter.LoadNetWordDict("https://raw.githubusercontent.com/Tohrusky/sensitive-go/main/dict/dict.txt")
```
## License

sensitive-go is licensed under the MIT License.
