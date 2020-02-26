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
    </div>
    <nav v-if="pagination">
      <ul class="pagination pagination-sm justify-content-end mt-2">
        <li class="page-item">
          <a class="page-link" href="javascript:void(0)" tabindex="-1" aria-label="Previous">上一页</a>
        </li>
        <li class="page-item active">
          <a class="page-link" href="javascript:void(0)">1</a>
        </li>
        <li class="page-item">
          <a class="page-link" href="javascript:void(0)">2</a>
        </li>
        <li class="page-item">
          <a class="page-link" href="javascript:void(0)">3</a>
        </li>
        <li class="page-item">
          <a class="page-link" href="javascript:void(0)">4</a>
        </li>
        <li class="page-item">
          <a class="page-link" href="javascript:void(0)" aria-label="Next">下一页</a>
        </li>
      </ul>
    </nav>
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