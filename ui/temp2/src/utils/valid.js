export const vType = {
  length (min, max, msg) {
    return {
      type: 'length',
      min,
      max,
      msg
    }
  },
  email (msg) {
    return {
      type: 'email',
      email: msg
    }
  },
  require (msg) {
    return {
      type: 'require',
      msg: msg
    }
  },
  regex (regx, msg) {
    return {
      type: 'regex',
      regx,
      msg
    }
  }
}

export const valid = {
  valid (item) {
    if (!item.valid) {
      item.validState = true
      return
    }
    let errCount = 0
    for (let key in item.valid) {
      const rule = item.valid[key]
      if (valid[rule.type]) {
        const flag = valid[rule.type](rule, item.value)
        if (flag) {
          errCount ++
          item.validState = false,
          item.validMsg = rule.msg
          return
        }
      } else {
        console.error('找不到验证方式：' + rule.type)
      }
    }
    if (errCount === 0) {
      item.validState = true,
      item.validMsg = ''
    }
  },
  require (rule, val) {
    return !val || val === ''
  },
  email (rule, val) {
    if (!val || val === '') {
      return false
    }
    const reg = /^\w+((.\w+)|(-\w+))@[A-Za-z0-9]+((.|-)[A-Za-z0-9]+).[A-Za-z0-9]+$/
    if (!reg.test(val)) {
      return true
    }
    return false
  },
  length (rule, val) {
    if (!val) {
      return false
    }
    const { max, min } = rule
    if (max && val.length > max) {
      return true
    }
    if (min && val.length < min) {
      return true
    }
    return false
  },
  regex (rule, val) {
    const {regx} = rule
    const re = new RegExp(regx, 'g')
    if (!re.test(val)) {
      return true
    }
    return false
  }
}

// valid.methods = {}

// valid.methods.email = function(elem) {
//   if (!elem.value || elem.value === '') {
//     return true
//   }
//   const reg = /^\w+((.\w+)|(-\w+))@[A-Za-z0-9]+((.|-)[A-Za-z0-9]+).[A-Za-z0-9]+$/
//   if (reg.test(elem.value)) {
//     return true
//   }
//   elem.valid = false
//   elem.msg = elem.email
//   return false
// }

// valid.methods.require = function(elem) {
//   if (!elem.value || elem.value === '') {
//     elem.valid = false
//     elem.msg = elem.require
//   } else {
//     elem.valid = true
//     elem.msg = ''
//   }
//   return elem.valid
// }

// valid.methods["length"] = function(elem) {
//   if (!elem.value) {
//     elem.valid = false
//   }
//   const { max, min, msg } = elem['length']
//   if (max && elem.value.length > max) {
//     elem.valid = false
//     elem.msg = msg
//   }
//   if (min && elem.value.length < min) {
//     elem.valid = false
//     elem.msg = msg
//   }
//   return elem.valid
// }

// valid.methods.regex = function(elem) {
//   const {reg, msg} = elem.regex
//   const re = new RegExp(reg, 'g')
//   if (re.test(elem.value)) {
//     return true
//   }
//   elem.valid = false
//   elem.msg = msg
// }

// valid.checkElement = function(key, elem) {
//   for (let type in elem) {
//     if (type === 'require' && elem[type] && !this.methods.require(elem)) {
//       return false
//     }
//     if (type === 'email' && elem[type] && !this.methods.email(elem)) {
//       return false
//     }
//     if (type === 'length' && elem[type] && !this.methods[type](elem)) {
//       return false
//     }
//     if (type === 'regex' && elem[type] && !this.methods.regex(elem)) {
//       return false
//     }
//   }
//   return true
// }

// valid.check = function(tmpValues) {
//   const values = JSON.parse(JSON.stringify(tmpValues))
//   let valid = true
//   for (let key in values) {
//     if (!this.checkElement(key, values[key])) {
//       valid = false
//     }
//   }
//   return {
//     values: values,
//     valid: valid
//   }
// }