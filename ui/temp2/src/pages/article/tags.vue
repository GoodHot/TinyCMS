<template>
  <t-block title="所有标签" theme rounded>
    <div class="push d-flex justify-content-between ">
      <div>
        <b-button size="sm" variant="light" class="mr-1" v-if="selectRows.length > 0"><t-icon icon="times" class="text-danger" /> 删除这{{ selectRows.length }}个标签</b-button>
        <b-button size="sm" variant="success" @click="$refs.tagModal.show()"><t-icon icon="plus" /> 创建标签</b-button>
      </div>
      <div>
        <b-input-group>
          <b-form-input placeholder="标签搜索" size="sm"></b-form-input>
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
      <template v-slot:path="prop">
        <b-badge>/tag/{{ prop.item.path }}</b-badge>
      </template>
      <template v-slot:action=" {item} = props">
        <b-button-group size="sm">
          <b-button variant="light" @click="deleteItem(item.id)"><t-icon icon="times" class="text-danger" /> 删除</b-button>
          <b-button variant="light" @click="$refs.tagModal.edit(item.id)"><t-icon icon="edit" class="text-primary" /> 编辑</b-button>
        </b-button-group>
      </template>
    </t-table>
    <tag-modal ref="tagModal" @onsave="onsaveHandler"></tag-modal>
  </t-block>
</template>
<script>
import { getTagByPage } from '@/api/tag'
import tagModal from './modal/tagModal'

export default {
  components: {
    tagModal
  },
  data () {
    return {
      title: '标签管理',
      description: '管理标签的展示和路径',
      nestableItems: [],
      columns: [
        {
          title: '图标',
          slot: 'icon',
          width: '90px'
        },
        {
          title: '标签名称',
          key: 'name',
          width: 150
        },
        {
          title: '标签路径',
          slot: 'path',
          width: 200
        },
        {
          title: '文章数',
          key: 'article_count',
          width: 100
        },
        {
          title: '描述',
          key: 'description'
        },
        {
          title: 'MetaTitle',
          key: 'meta_title'
        },
        {
          title: 'MetaDescription',
          key: 'meta_description'
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
        baseURL: '/#/tags/'
      },
      selectRows: [],
      tableLoading: false
    }
  },
  created () {
    this.loadTag()
  },
  methods: {
    loadTag () {
      this.tableLoading = true
      getTagByPage(this.pagination.current, '').then(res => {
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
      this.loadTag()
    },
    onsaveHandler () {
      alert(`保存标签成功`)
      this.loadTag()
    }
  }
}
</script>