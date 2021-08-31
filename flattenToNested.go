package flattenToNested
import(
    "strings"
)

func flattenToNested(a map[string]interface{}) map[string]interface{} {
    b := make(map[string]interface{}, len(a))
    for key, value := range a{
        res := strings.FieldsFunc(key, func (c rune) bool {
            if c == '.'{
                return true
            }
            return false
        })
        if len(res) > 1{
            parseDeep(b, res, value)
        } else{
            b[key] = value
        }
            
    }
   
    return b
}


func  parseDeep(a map[string]interface{}, keys []string, value interface{})  {
    b := a
    length := len(keys)
    for i := 0; i < length; i++ {
        k := keys[i]
        v,exist := b[k]
        if !exist {
            //k为键的映射不存在，则创建新映射或拓展原映射的空间以容纳新的键值对
            if i == length - 1{
                b[k] = value
            }else{
                var c map[string]interface{}
                _, ok := b[k].(map[string]interface{})
                if ok {  //由于b[k]指向已存在映射，因此拷贝源映射的值至新的映射
                    c = make(map[string]interface{}, len(b)+1)
                    for k1,v1 := range b{
                        c[k1] = v1
                        delete(b,k1)
                    }
                } else {//由于b[k]不指向映射，单纯创建新映射内存空间
                    c = make(map[string]interface{}, 1)
                }
                b[k] = c
                b = c  
            }         
        } else {
            //k为键的映射已经存在，如果b[k]是一个映射，则b指向它
            n, ok := v.(map[string]interface{})
            if ok {
                b = n       
            }               
        }
    }
}