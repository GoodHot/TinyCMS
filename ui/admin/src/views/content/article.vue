<template>
  <a-card :bordered="false">
    <a-table
      :columns="columns"
      :dataSource="data"
      rowKey="id"
      :pagination="pagination"
      :loading="loading"
      @change="handleTableChange"
    >
      <template slot="action" slot-scope="text, record">
        <a href="" v-if="record.status == 2">发布</a>
        <a href="" v-if="record.status == 1">下架</a>
        <a-divider type="vertical" />
        <a @click="handlerEdit(record)">编辑</a>
        <a-divider type="vertical" />
        <a-popconfirm
          title="确定要删除这篇文章?"
          @confirm="handleDelete(record)"
          okText="删除"
          cancelText="取消"
        >
          <a href="#">删除</a>
        </a-popconfirm>
      </template>
      <template slot="status" slot-scope="text, record">
        <a-tag color="#87d068" v-if="record.status == 1">发布</a-tag>
        <a-tag color="#2db7f5" v-if="record.status == 2">下架</a-tag>
      </template>
    </a-table>
  </a-card>
</template>
<script>
import { getArticlePage, deleteArticle } from '@/api/article'
export default {
  data () {
    return {
      columns: [
        {
          title: '#',
          width: 120,
          dataIndex: 'id'
        },
        {
          title: '标题',
          width: 200,
          dataIndex: 'title'
        },
        {
          title: '频道名称',
          width: 150,
          dataIndex: 'channel_id'
        },
        {
          title: '状态',
          width: 50,
          scopedSlots: { customRender: 'status' }
        },
        {
          title: '描述',
          dataIndex: 'description'
        },
        {
          title: '浏览数',
          width: 100,
          dataIndex: 'views'
        },
        {
          title: '操作',
          width: 200,
          scopedSlots: { customRender: 'action' }
        }
      ],
      data: [],
      loading: false,
      pagination: {},
      param: {
        page: 1
      }
    }
  },
  mounted () {
    this.loadData()
  },
  methods: {
    loadData () {
      this.loading = true
      getArticlePage(this.param).then(res => {
        const result = res.page
        const pagination = { ...this.pagination }
        pagination.total = result.total_count
        pagination.pageSize = result.page_size
        pagination.current = result.page_num
        this.pagination = pagination
        this.data = result.list
        this.loading = false
      })
    },
    handleTableChange (pagination, filters, sorter) {
      this.param.page = pagination.current
      this.loadData()
    },
    handleDelete (data) {
      deleteArticle(data.id).then(res => {
        this.$message.success(`操作成功`)
        this.loadData()
      })
    },
    handlerEdit (data) {
      this.$router.push(`/content/publish?article=${data.id}`)
    }
  }
}
</script>
