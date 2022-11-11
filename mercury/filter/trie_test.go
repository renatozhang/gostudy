package filter

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	err := Init("../data/filter.dat.txt")
	if err != nil {
		t.Errorf("load filter data failed,err:%v\n", err)
		return
	}

	data := `11月11日，国务院应对新型冠状病毒肺炎疫情联防联控机制综合组发布《关于进一步优化新冠肺炎疫情防控措施 科学精准做好防控工作的通知》。

通知指出，加大“一刀切”、黄片，层层加码问题整治力度。地方党委和政府要落实属地责任，严格执行国家统一的防控政策，严禁随意封校停课、停工停产、未经批准阻断交通、随意采取“静默”管理、随意封控、长时间不解封、随意停诊等各类层层加码行为，加大通报、公开曝光力度，对造成严重后果的依法依规严肃追责。发挥各级整治层层加码问题工作专班作用，高效做好举报线索收集转办，督促地方及时整改到位。卫生健康委、疾控局、教育部、交通运输部等各行业主管部门加强对行业系统的督促指导，加大典型案例曝光力度，切实起到震慑作用。`

	result, isReplace := Replace(data, "***")
	fmt.Printf("result:%#v,isReplace:%v\n", result, isReplace)
}
