const valid = {}

valid.methods = {}

valid.methods.email = function(elem) {
  if (!elem.value || elem.value === '') {
    return true
  }
  const reg = /^\w+((.\w+)|(-\w+))@[A-Za-z0-9]+((.|-)[A-Za-z0-9]+).[A-Za-z0-9]+$/
  if (reg.test(elem.value)) {
    return true
  }
  elem.valid = false
  elem.msg = elem.email
  return false
}

valid.methods.require = function(elem) {
  if (!elem.value || elem.value === '') {
    elem.valid = false
    elem.msg = elem.require
  } else {
    elem.valid = true
    elem.msg = ''
  }
  return elem.valid
}

valid.methods["length"] = function(elem) {
  if (!elem.value) {
    elem.valid = false
  }
  const { max, min, msg } = elem['length']
  if (max && elem.value.length > max) {
    elem.valid = false
    elem.msg = msg
  }
  if (min && elem.value.length < min) {
    elem.valid = false
    elem.msg = msg
  }
  return elem.valid
}

valid.checkElement = function(key, elem) {
  for (let type in elem) {
    if (type === 'require' && elem[type] && !this.methods.require(elem)) {
      return false
    }
    if (type === 'email' && elem[type] && !this.methods.email(elem)) {
      return false
    }
    if (type === 'length' && elem[type] && !this.methods[type](elem)) {
      return false
    }
  }
  return true
}

valid.check = function(tmpValues) {
  const values = JSON.parse(JSON.stringify(tmpValues))
  let valid = true
  for (let key in values) {
    if (!this.checkElement(key, values[key])) {
      valid = false
    }
  }
  return {
    values: values,
    valid: valid
  }
}

export default valid