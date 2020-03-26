<template>
  <t-block title="所有分类" theme rounded>
    <div class="push d-flex justify-content-between ">
      <div>
        <b-button size="sm" variant="light" class="mr-1" v-if="selectRows.length > 0"><t-icon icon="times" class="text-danger" /> 删除这{{ selectRows.length }}个分类</b-button>
        <b-button size="sm" class="mr-1" @click="$refs.categorySortModal.show()"><t-icon icon="sort-amount-down" /> 设置排序</b-button>
        <b-button size="sm" variant="success"><t-icon icon="plus" /> 创建分类</b-button>
      </div>
      <div>
        <b-input-group>
          <b-form-input placeholder="分类搜索" size="sm"></b-form-input>
          <b-input-group-append>
            <b-button size="sm" variant="dark" class="px-3"><t-icon size="fw" icon="search" /></b-button>
          </b-input-group-append>
        </b-input-group>
      </div>
    </div>
    <t-table ref="dataTable" :loading="tableLoading" @onpage="onpageHandler" @onselect="onselectHandler" selectRow rowKey="id" :columns="columns" :dataSource="data" :pagination="pagination" striped hover>
      <template v-slot:icon="prop">
        <b-img :src="assetsURL(prop.item.icon)" :width="32" :height="32" rounded alt="Rounded image"></b-img>
      </template>
      <template v-slot:action=" {item} = props">
        <b-button-group size="sm">
          <b-button variant="light" @click="deleteItem(item.id)"><t-icon icon="times" class="text-danger" /> 删除</b-button>
          <b-button variant="light"><t-icon icon="edit" class="text-primary" /> 编辑</b-button>
        </b-button-group>
      </template>
    </t-table>
    <category-sort-modal ref="categorySortModal"></category-sort-modal>
  </t-block>
</template>
<script>
import { getCategoryByPage } from '@/api/category'
import CategorySortModal from './modal/categorySort'

export default {
  components: {
    CategorySortModal
  },
  data () {
    return {
      title: '分类和标签管理',
      description: '添加分类和查看标签，可以通过拖拽移动分类排序',
      nestableItems: [],
      columns: [
        {
          title: '图标',
          slot: 'icon',
          width: '90px'
        },
        {
          title: '分类名称',
          key: 'name',
          width: 200
        },
        {
          title: '路径',
          key: 'path',
          width: 200
        },
        {
          title: '文章数',
          key: 'article_count',
          width: 100
        },
        {
          title: '备注',
          key: 'remark'
        },
        {
          title: '操作',
          slot: 'action',
          width: 160
        }
      ],
      data: [],
      pagination: {
        total: 1,
        current: 1,
        baseURL: '/#/category/'
      },
      selectRows: [],
      tableLoading: false
    }
  },
  created () {
    this.loadCategory()
  },
  methods: {
    loadCategory () {
      this.tableLoading = true
      getCategoryByPage(this.pagination.current, '').then(res => {
        this.data = res.page.list,
        this.pagination.total = res.page.total_page
        this.tableLoading = false
      })
    },
    onselectHandler (all, items) {
      this.selectRows = items
    },
    onpageHandler (page) {
      this.pagination.current = page
      this.loadCategory()
    }
  }
}
</script>