package dir

import "additive-apps/additive-lab/src/util"

func NotExists(path string, operations ...func()) bool {
	return util.ResourceNotExists("Directory", path, operations...)
}
