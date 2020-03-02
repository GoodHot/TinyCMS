<template>
  <TTable :column="column" :pagination="false" :data="data" class="table table-hover table-vcenter font-size-sm">
    <template v-slot:options="{item}">
      <b-form-checkbox
        v-for="option in item.values"
        :key="option.value"
        inline
      >
        {{ option.text }}
      </b-form-checkbox>
    </template>
  </TTable>
</template>
<script>
import {allPermission} from '@/api/role'

export default {
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
      selected: [], // Must be an array reference!
      data: [
        {
          text: "文章管理",
          values: [
            { text: '全选', value: 'all' },
            { text: '发布文章', value: 'orange' },
            { text: '编辑文章', value: 'apple' },
            { text: '删除文章', value: 'pineapple' },
            { text: '查询文章列表', value: '1' }
          ]
        },
        {
          text: "分类管理",
          values: [
            { text: '全选', value: 'all' },
            { text: '编辑分类', value: '2' },
            { text: '删除分类', value: '3' },
            { text: '增加分类', value: '4' },
            { text: '查询分类', value: '0' }
          ]
        },
        {
          text: "用户管理",
          values: [
            { text: '全选', value: 'all' },
            { text: '发布文章', value: 'orange' },
            { text: '编辑文章', value: 'apple' },
            { text: '删除文章', value: 'pineapple' },
            { text: '查询文章列表', value: '1' }
          ]
        },
        {
          text: "用户管理",
          values: [
            { text: '全选', value: 'all' },
            { text: '编辑分类', value: '2' },
            { text: '删除分类', value: '3' },
            { text: '增加分类', value: '4' }
          ]
        }
      ]
    }
  },
  mounted() {
    this.loadPermission()
  },
  methods: {
    loadPermission() {
      allPermission().then(res => {
        const pers = []
        for (let i in res.permission) {
          const p = res.permission[i]
          const child = [
            {
              text: '全选',
              value: 'all'
            }
          ]
          for (let j in p.child) {
            const c = p.child[j]
            child.push(
              {
                text: c.permission_name,
                value: c.id
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
    }
  }
}
</script>