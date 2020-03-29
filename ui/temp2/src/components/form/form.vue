<template>
  <b-form>
    <b-form-group :label="item.label" v-for="(item, index) of items" :key="index">
      <b-form-input v-if="item.type === 'input'" @blur="validation(item)" :placeholder="item.placeholder" v-model="item.value" :state="item.validState"></b-form-input>
      <b-form-textarea v-if="item.type === 'textarea'" @blur="validation(item)" rows="3" max-rows="6" v-model="item.value" :placeholder="item.placeholder" :state="item.validState"></b-form-textarea>
      <t-image-upload v-if="item.type === 'imageUpload'" v-model="item.value" :state="item.validState"></t-image-upload>
      <b-form-invalid-feedback :state="item.validState">
        {{item.validMsg}}
      </b-form-invalid-feedback>
    </b-form-group>
  </b-form>
</template>
<script>
import {PropTypes} from '@/utils/types'
import { valid } from '@/utils/valid'

export default {
  props: {
    form: PropTypes.Array
  },
  data () {
    return {
      items: null
    }
  },
  mounted () {
    const temp = JSON.parse(JSON.stringify(this.form))
    temp.map(item => {
      item.validState = null
      item.validMsg = ''
    })
    this.items = temp
  },
  methods: {
    validation (item) {
      valid.valid(item)
    },
    validate () {
      let hasErr = false
      this.items.map(item => {
        valid.valid(item)
        if (!item.validState) {
          hasErr = true
        }
      })
      return !hasErr
    },
    submit (callback) {
      const data = {}
      this.items.map(item => {
        data[item.name] = item.value
      })
      callback(data)
    }
  }
}
</script>