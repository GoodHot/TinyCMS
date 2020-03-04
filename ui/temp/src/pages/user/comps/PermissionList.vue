<template>
  <TTable :column="column" :pagination="false" :data="data" class="table table-hover table-vcenter font-size-sm">
    <template v-slot:options="{item}">
      <b-form-checkbox
        v-for="option in item.values"
        :key="option.value"
        v-model="option.checked"
        inline
        @change="checkboxChangeHandler(option, item.values)"
        :disabled="isSuper"
      >
        {{ option.text }}
      </b-form-checkbox>
    </template>
  </TTable>
</template>
<script>
import {allPermission} from '@/api/role'

export default {
  props: {
    checkedValues: {
      type: Array,
      default() {
        return []
      }
    },
    isSuper: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      column: [
        {
          title: "权限",
          index: "text",
          width: "120px",
          class: "d-none d-sm-table-cell font-w600"
        },
        {
          title: "选项",
          index: "values",
          slot: "options"
        }
      ],
      data: []
    }
  },
  mounted() {
    this.loadPermission()
  },
  methods: {
    checkboxChangeHandler(option, values) {
      option.checked = !option.checked
      if (option.value === 'all') {
        for (let i=1;i<values.length;i++) {
          values[i].checked = option.checked
        }
        return
      }
      for (let i=1;i<values.length;i++) {
        if (!values[i].checked) {
          values[0].checked = false
          return 
        }
      }
      values[0].checked = true
    },
    loadPermission() {
      allPermission().then(res => {
        const pers = []
        for (let i in res.permission) {
          const p = res.permission[i]
          const child = [
            {
              text: '全选',
              value: 'all',
              checked: false
            }
          ]
          for (let j in p.child) {
            const c = p.child[j]
            child.push(
              {
                text: c.permission_name,
                value: c.id,
                checked: false
              }
            )
          }
          pers.push(
            {
              text: p.permission_name,
              values: child
            }
          )
        }
        this.data = pers
      })
    },
    getValue() {
      const ids = []
      for (let i in this.data) {
        for (let j in this.data[i].values) {
          const per = this.data[i].values[j]
          if (per.checked && per.value !== 'all') {
            ids.push(per.value)
          }
        }
      }
      return ids
    }
  },
  watch: {
    checkedValues(val) {
      console.log('ck', val)
      console.log('dt', this.data)
      for (let i in this.data) {
        for (let j in this.data[i].values) {
          for (let k in val) {
            if (this.data[i].values[j].value === val[k]) {
              this.data[i].values[j].checked = true
            }
          }
        }
      }
    }
  }
}
</script>