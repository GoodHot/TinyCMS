<template>
  <t-block title="系统设置" border theme>
    <b-form>
      <b-row class="push">
        <b-col :cols="4">
          <div v-for="dict in dicts" :key="dict.id">
            <b-form-group
              :description="dict.description"
              :label="dict.name"
              v-if="dict.visible"
            >
              <b-form-input v-if="dict.type === 'text'" v-model="dict.value"></b-form-input>
              <b-form-textarea v-if="dict.type === 'textarea'" v-model="dict.value" rows="3" max-rows="6"></b-form-textarea>
              <b-form-select v-if="dict.type === 'select'" v-model="dict.value" :options="getItemOptions(dict)"></b-form-select>
            </b-form-group>
          </div>
          <b-button variant="primary" @click="save">保存设置</b-button>
        </b-col>
        <b-col :cols="8"></b-col>
      </b-row>
    </b-form>
  </t-block>
</template>
<script>
import { getAllDict, saveDict } from '@/api/dict'
export default {
  data () {
    return {
      title: '系统设置',
      description: '可配置全局系统',
      dicts: []
    }
  },
  mounted () {
    getAllDict().then(res => {
      this.dicts = res.dicts
    })
  },
  methods: {
    getItemOptions (dict) {
      const opts = dict.options.split(',')
      const rst = []
      opts.map(opt => {
        const temp = opt.split(':')
        rst.push({
          value: temp[1],
          text: temp[0]
        })
      })
      return rst
    },
    save () {
      saveDict({
        dicts: this.dicts
      }).then(res => {
        console.log(res)
      })
    }
  }
}
</script>