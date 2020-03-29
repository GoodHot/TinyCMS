<template>
  <b-overlay :show="loading" rounded="sm">
    <table :class="classes">
      <thead>
        <tr>
          <th v-if="selectRow" width="20">
            <b-form-checkbox v-model="isSelectAll" @change="selectAll"></b-form-checkbox>
          </th>
          <th v-for="(col, index) in columns" :key="index" :width="col.width ? col.width: 'auto'">{{ col.title }}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, i) in items" :key="i">
          <td v-if="selectRow">
            <b-form-checkbox :ref="selectRow" @change="selectOne($event, item)" v-model="item.checked"></b-form-checkbox>
          </td>
          <td v-for="(col, index) in columns" :key="index">
            <span v-if="!col.slot">{{ item[col.key] }}</span>
            <slot v-else :name="col.slot" v-bind:item="item" v-bind:index="i"></slot>
          </td>
        </tr>
      </tbody>
    </table>
    <b-row>
      <b-col :cols="6">
        Page 1 of 4
      </b-col>
      <b-col :cols="6" class="text-right">
        <b-pagination-nav @change="pageChangeHandler" :base-url="pagination.baseURL" v-model="pagination.current" class="d-inline-block" v-if="pagination" :number-of-pages="pagination.total"></b-pagination-nav>
      </b-col>
    </b-row>
  </b-overlay>
</template>
<script>
import {PropTypes} from '@/utils/types'

export default {
  props: {
    columns: PropTypes.Array,
    dataSource: PropTypes.Array,
    size: PropTypes.String, // sm, lg
    pagination: PropTypes.Object,
    striped: PropTypes.Boolean.def(false),
    hover: PropTypes.Boolean.def(false),
    bordered: PropTypes.Boolean.def(false),
    borderless: PropTypes.Boolean.def(false),
    rowKey: PropTypes.String.def('id'),
    selectRow: PropTypes.Boolean.def(false),
    loading: PropTypes.Boolean.def(false)
  },
  data () {
    return {
      items: [],
      isSelectAll: false
    }
  },
  computed: {
    classes () {
      return [
        'table',
        'table-vcenter',
        {
          'table-striped': this.striped,
          'table-hover': this.hover,
          'table-bordered': this.bordered,
          'table-borderless': this.borderless
        }
      ]
    }
  },
  created () {
    this.refreshData(this.dataSource)
  },
  methods: {
    selectAll (checked) {
      this.isSelectAll = checked
      this.items.map(item => {
        item.checked = checked
      })
      this.onSelect()
    },
    selectOne (checked, item) {
      item.checked = checked
      let checkedCount = 0
      this.items.map(item => {
        if (!item.checked) {
          this.isSelectAll = false
        } else {
          checkedCount ++
        }
      })
      if (checkedCount === this.items.length ) {
        this.isSelectAll = true
      }
      this.onSelect()
    },
    onSelect () {
      const rows = []
      this.items.map(item => {
        if (item.checked) {
          rows.push(item)
        }
      })
      this.$emit('onselect', this.isSelectAll, rows)
    },
    refreshData (val) {
      this.isSelectAll = false
      const items = JSON.parse(JSON.stringify(val))
      items.map(item => {
        item.checked = false
      })
      this.items = items
      this.$emit('onselect', this.isSelectAll, [])
    },
    pageChangeHandler (page) {
      this.$emit('onpage', page)
    }
  },
  watch: {
    dataSource (val) {
      this.refreshData(val)
    }
  }
}
</script>