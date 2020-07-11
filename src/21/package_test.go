package _21_test

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

/**
1,基本复用模块但愿
	以首字母大写来表明可以被包外代码访问
2，代码的package可以和所在的项目不一致
3，统一目录的Go代码的package要保持一致
*/

/**
1，在main被执行前，所有依赖的package的init方法都会被执行
2，不同包的init函数按照包倒入的依赖关系决定执行吮吸
3，每个包可以有多个init函数
4，包的每个源文件也可以有过个init函数。


*/

/**
操作步骤：  https://blog.csdn.net/sinat_28545681/article/details/52535720?utm_medium=distribute.pc_aggpage_search_result.none-task-blog-2~all~first_rank_v2~rank_v25-1-52535720.nonecase

 */

func TestUseDependencyFromGit(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"),"value-111")
	t.Log(m.Get(cm.StrKey("key")))
}
