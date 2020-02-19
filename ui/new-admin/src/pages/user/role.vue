<template>
  <TBlock title="角色管理">
    <div class="pull-x">
      <TTable :pagination="true" :column="column" :data="data" ref="dataTable" class="table table-hover table-vcenter font-size-sm">
        <template v-slot:roleName="{item}" class="d-none d-sm-table-cell font-w600">
          {{item.role_name}}
        </template>
        <template v-slot:permission="{item}">
          <div v-for="per of item.permission" :key="per.code" class="mb-2">
            <TBadge type="primary">{{per.permission_name}}</TBadge>
            <TBadge class="ml-2" type="info" v-for="cl of per.children" :key="cl.code">{{cl.permission_name}}</TBadge>
          </div>
        </template>
        <template v-slot:action>
          <TButtonGroup size="sm" type="light" v-slot:default="{size, type, theme}" class="mr-1" >
              <TButton :size="size" :type="type" :theme="theme" icon="edit" iconPack="far">编辑</TButton>
              <TButton :size="size" :type="type" :theme="theme" icon="trash-alt" iconPack="far">删除</TButton>
            </TButtonGroup>
        </template>
      </TTable>
    </div>
  </TBlock>
</template>
<script>
export default {
  mounted() {
  },
  data() {
    return {
      column: [
        {
          title: 'ID',
          index: 'id',
          align: 'center',
          width: '100px',
          class: 'd-none d-sm-table-cell font-w600'
        },
        {
          title: '编码',
          index: 'code',
          width: '150px',
          class: 'd-none d-sm-table-cell'
        },
        {
          title: '角色',
          index: 'role_name',
          slot: 'roleName',
          width: '250px',
          class: 'font-w600'
        },
        {
          title: '权限',
          slot: 'permission'
        },
        {
          title: '用户数',
          index: 'admin_count',
          width: '120px',
          class: 'd-none d-sm-table-cell font-w600'
        },
        {
          title: '操作',
          slot: 'action',
          width: '200px',
          class: 'd-none d-xl-table-cell text-muted'
        }
      ],
      data: [
        {
          id: 123,
          code: 'super',
          role_name: '超级管理员',
          admin_count: 123,
          permission: [
            {
              code: 'content',
              permission_name: '内容管理',
              children: [
                {
                  code: 'content_publish',
                  permission_name: '发布文章'
                },
                {
                  code: 'content_edit',
                  permission_name: '编辑文章'
                },
                {
                  code: 'content_delete',
                  permission_name: '删除文章'
                },
                {
                  code: 'content_list',
                  permission_name: '查询文章列表'
                }
              ]
            },
            {
              code: 'user',
              permission_name: '用户管理',
              children: [
                {
                  code: 'user_edit',
                  permission_name: '编辑用户'
                },
                {
                  code: 'user_delete',
                  permission_name: '删除用户'
                },
                {
                  code: 'user_list',
                  permission_name: '查询用户列表'
                }
              ]
            }
          ]
        },
        {
          id: 125,
          code: 'super',
          role_name: '超级管理员',
          admin_count: 123,
          permission: [
            {
              code: 'content',
              permission_name: '内容管理',
              children: [
                {
                  code: 'content_publish',
                  permission_name: '发布文章'
                },
                {
                  code: 'content_edit',
                  permission_name: '编辑文章'
                },
                {
                  code: 'content_delete',
                  permission_name: '删除文章'
                },
                {
                  code: 'content_list',
                  permission_name: '查询文章列表'
                }
              ]
            },
            {
              code: 'user',
              permission_name: '用户管理',
              children: [
                {
                  code: 'user_edit',
                  permission_name: '编辑用户'
                },
                {
                  code: 'user_delete',
                  permission_name: '删除用户'
                },
                {
                  code: 'user_list',
                  permission_name: '查询用户列表'
                }
              ]
            }
          ]
        }
      ]
    }
  }
}
</script>