<template>
  <a-card :bordered="false">
    <div class="table-operator">
      <a-button type="primary" icon="plus" @click="$refs.editChannel.loadAdd(0)" >新建频道</a-button>
    </div>
    <a-divider />
    <a-table
      :columns="columns"
      :dataSource="data"
      rowKey="id"
      :pagination="pagination"
      :loading="loading"
      @change="handleTableChange"
    >
      <template slot="action" slot-scope="text, record">
        <a @click="$refs.editChannel.loadAdd(record.id)">创建子频道</a>
        <a-divider type="vertical" />
        <a @click="handlerEdit(record)">编辑</a>
        <a-divider type="vertical" />
        <a @click="handleDelete(record)">删除</a>
      </template>
      <template slot="sort">
        <a-button-group size="small">
          <a-button type="dashed" icon="up" />
          <a-button type="dashed" icon="down" />
        </a-button-group>
      </template>
      <template slot="icon" slot-scope="text, record">
        <span v-html="record.icon"></span>
      </template>
    </a-table>
    <edit-channel ref="editChannel" @commitHandler="commitHandler"></edit-channel>
  </a-card>
</template>
<script>
import EditChannel from './module/EditChannel'
import { getChannelPage, deleteChannel } from '@/api/channel'
export default {
  components: {
    EditChannel
  },
  data () {
    return {
      columns: [
        {
          title: '#',
          width: 120,
          dataIndex: 'id'
        },
        {
          title: '图标',
          width: 100,
          scopedSlots: { customRender: 'icon' }
        },
        {
          title: '频道名称',
          width: 230,
          dataIndex: 'name'
        },
        {
          title: '频道英文',
          width: 230,
          dataIndex: 'title'
        },
        {
          title: '排序',
          width: 130,
          scopedSlots: { customRender: 'sort' }
        },
        {
          title: '备注',
          dataIndex: 'remark'
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
        page: 1,
        parentId: 0
      }
    }
  },
  mounted () {
    this.loadData()
  },
  methods: {
    commitHandler () {
      this.loadData()
    },
    loadChildData (record) {
      getChannelPage({
        parentId: record.id,
        page: 1
      }).then(res => {
        record.children = res.page.list
      })
    },
    loadData () {
      this.loading = true
      getChannelPage(this.param).then(res => {
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
      deleteChannel(data.id).then(res => {
        this.$message.success(`操作成功`)
        this.loadData()
      })
    },
    handlerEdit (data) {
      this.$refs.editChannel.loadEdit(data.id)
    }
  }
}
</script>
