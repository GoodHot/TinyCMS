<template>
  <div>
    <div class="table-responsive">
      <table class="table table-borderless table-striped table-vcenter">
        <thead>
          <tr>
            <th :class="thClass(item)" v-for="item of column" :key="item.title">{{item.title}}</th>
            <!-- <th class="d-none d-sm-table-cell text-center">Submitted</th>
            <th>Status</th>
            <th class="d-none d-xl-table-cell">Customer</th>
            <th class="d-none d-xl-table-cell text-center">Products</th>
            <th class="d-none d-sm-table-cell text-right">Value</th>
            <th class="text-center">Action</th> -->
          </tr>
        </thead>
        <tbody>
          <tr v-for="(d, index) of data" :key="d.id">
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
      <ul class="pagination justify-content-end mt-2">
        <li class="page-item">
          <a class="page-link" href="javascript:void(0)" tabindex="-1" aria-label="Previous">Prev</a>
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
          <a class="page-link" href="javascript:void(0)" aria-label="Next">Next</a>
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
    }
  },
  methods: {
    thClass(col) {
      let cls = ''
      if (col.align && col.align == 'center') {
        cls = 'text-center '
      }
      return cls
    }
  }
}
</script>