<template>
  <div>
    <div class="table-responsive">
      <table :class="'table table-borderless table-striped table-vcenter'">
        <thead v-if="!hideHeader">
          <tr>
            <th v-if="selectKey" width="60px" class="text-center">
              <TCheckbox type="secondary" name="selectAll" v-model="selectAllValue" @change="$selectAll" />
              <b-form-checkbox
                name="selectAll"
              >
                I accept the terms and use
              </b-form-checkbox>
            </th>
            <th :class="thClass(item)" v-for="item of column" :key="item.title">{{item.title}}</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(d, index) of data" :key="d.id">
            <td v-if="selectKey" width="60px" class="text-center">
              <b-form-checkbox :name="`select-${index}`" v-model="selectValues[index].checked" :ref="`selectItems`" @change="selectOne"></b-form-checkbox>
            </td>
            <template v-for="col of column">
              <td :class="thClass(col)" :key="col.title" :width="col.width ? col.width : ''">
                <div v-if="!col.slot">
                  {{ d[col.index] }}
                </div>
                <slot v-else :name="col.slot" v-bind:item="d" v-bind:index="index"></slot>
              </td>
            </template>
          </tr>
        </tbody>
      </table>
      <div class="alert alert-secondary alert-dismissable text-center" v-if="!data || data.length === 0">
          <h3 class="alert-heading font-w300 my-2">暂无数据</h3>
          <b-button size="sm" variant="success">发布一篇文章</b-button>
      </div>
    </div>
    <!-- <Tpagination v-if="pagination" :data="pagination"></Tpagination> -->
    <b-pagination
      class="justify-content-center"
      v-model="pagination.pageNum"
      :total-rows="pagination.totalPage"
      per-page="1"
      first-text="首页"
      prev-text="上一页"
      next-text="下一页"
      last-text="尾页"
      @change="pageChangeHandler"
    ></b-pagination>
  </div>
</template>
<script>
export default {
  props: {
    column: {
      type: Array
    },
    data: {
      type: Array,
      default() {
        return []
      }
    },
    pagination: {
      type: [Boolean, Object],
      default: false
    },
    selectKey: {
      type: String
    },
    hideHeader: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      selectAllValue: false,
      selectValues: []
    }
  },
  mounted() {
    this.initSelectValues()
  },
  methods: {
    pageChangeHandler(page) {
      this.$emit('onpage', page)
    },
    initSelectValues() {
      if (!this.selectKey || !this.data) {
        return 
      }
      for (let i in this.data) {
        this.selectValues.push({
          index: i,
          checked: false,
          value: this.data[i][this.selectKey]
        })
      }
    },
    thClass(col) {
      let cls = ' '
      if (col.align && col.align == 'center') {
        cls += 'text-center '
      }
      if (col.class) {
        cls += col.class + ' '
      }
      return cls
    },
    $selectAll(rst) {
      this.selectAll(rst.checked)
    },
    selectAll(checked) {
      for (let i in this.selectValues) {
        this.selectValues[i].checked = checked
      }
      this.onSelected()
    },
    selectOne(rst) {
      this.onSelected()
      if (!rst.checked) {
        this.selectAllValue = false
      }
      for (let i in this.selectValues) {
        if (!this.selectValues[i].checked) {
          return 
        }
      }
      this.selectAllValue = true
    },
    onSelected() {
      const values = new Array()
      for (let i in this.selectValues) {
        if (this.selectValues[i].checked) {
          values.push(this.selectValues[i].value)
        }
      }
      this.$emit("onselected", values, values.length == this.selectValues.length)
    }
  },
  watch: {
    data() {
      this.initSelectValues()
    }
  }
}
</script>