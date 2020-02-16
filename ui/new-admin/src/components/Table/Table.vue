<template>
  <div>
    <div class="table-responsive">
      <table :class="'table table-borderless table-striped table-vcenter'">
        <thead v-if="!hideHeader">
          <tr>
            <th v-if="select" width="60px"><TCheckbox type="secondary" name="selectAll" ref="selectAll" @change="selectAll" /></th>
            <th :class="thClass(item)" v-for="item of column" :key="item.title">{{item.title}}</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(d, index) of data" :key="d.id">
            <td v-if="select" width="60px" class="text-center"><TCheckbox type="secondary" :name="`select-${index}`" :ref="`select${index}`" @change="selectOne" /></td>
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
    <nav v-if="pagination.show">
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
      type: Object,
      default() {
        return {
          show: false
        }
      }
    },
    select: {
      type: Boolean,
      default: false
    },
    hideHeader: {
      type: Boolean,
      default: false
    }
  },
  methods: {
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
    selectAll(data) {
      for (let i=0;i<this.data.length;i++) {
        const sel = this.$refs[`select${i}`]
        sel[0].check(data.checked)
      }
    },
    selectOne(data) {
      if (!data.checked) {
        if (this.$refs.selectAll) {
          this.$refs.selectAll.check(false)
        }
      }
      for (let i=0;i<this.data.length;i++) {
        const sel = this.$refs[`select${i}`]
        if (!sel[0].state()) {
          return
        }
      }
      if (this.$refs.selectAll) {
        this.$refs.selectAll.check(true)
      }
    }
  }
}
</script>