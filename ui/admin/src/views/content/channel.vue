<template>
  <a-card :bordered="false">
    <div class="table-operator">
      <a-button type="primary" icon="plus" @click="$refs.editChannel.loadAdd()" >新建频道</a-button>
    </div>
    <a-divider />
    <a-table
      :columns="columns"
      :dataSource="data"
      rowKey="id"
      :pagination="false"
      :loading="loading"
    >
      <template slot="action" slot-scope="text, record">
        <a @click="edit(record)">编辑</a>
        <a-divider type="vertical" />
        <a @click="handleDelete(record)">删除</a>
      </template>
    </a-table>
    <edit-channel ref="editChannel" @commitHandler="commitHandler"></edit-channel>
  </a-card>
</template>
<script>
import EditChannel from './module/EditChannel'
export default {
  components: {
    EditChannel
  },
  data () {
    return {
      columns: [
        {
          title: '#',
          width: 20,
          dataIndex: 'id'
        },
        {
          title: '频道名称',
          width: 320,
          dataIndex: 'name'
        },
        {
          title: '频道英文',
          width: 320,
          dataIndex: 'link'
        },
        {
          title: '排序',
          width: 100,
          dataIndex: 'sort'
        },
        {
          title: '备注',
          dataIndex: 'sort'
        },
        {
          title: '操作',
          width: 130,
          scopedSlots: { customRender: 'action' }
        }
      ],
      data: [],
      loading: false
    }
  },
  methods: {
    commitHandler () {
    }
  }
}
</script>
