<template>
  <div class="row">
    <div class="col-md-5 col-xl-3">
      <div class="d-md-none push">
        <button type="button" class="btn btn-block btn-primary" @click="showCategory">文章分类</button>
      </div>
      <div ref="category" class="d-none d-md-block push">
        <TBlock title="文章分类" headerBg>
          <template slot="options">
            <TButton icon="plus"></TButton>
            <TButton icon="settings" iconPack="si"></TButton>
          </template>
          <ul class="nav nav-pills flex-column font-size-sm push">
            <li class="nav-item my-1">
              <a class="nav-link d-flex justify-content-between align-items-center" href>
                <span>
                  <i class="fa fa-fw fa-book mr-1"></i> 全部文章
                </span>
                <span class="badge badge-pill badge-secondary">3篇</span>
              </a>
            </li>
            <li class="nav-item my-1">
              <a class="nav-link d-flex justify-content-between align-items-center" href>
                <span>
                  <i class="fa fa-fw fa-inbox mr-1"></i> 未归档
                </span>
                <span class="badge badge-pill badge-secondary">3篇</span>
              </a>
            </li>
          </ul>
        </TBlock>
        <TBlock title="标签" icon="tags">
          <ul class="nav nav-pills flex-column push">
            <li class="nav-item mb-1">
              <a
                class="nav-link d-flex justify-content-between align-items-center"
                href="javascript:void(0)"
              >
                npm
                <span class="badge badge-pill badge-secondary ml-1">3篇</span>
              </a>
            </li>
            <li class="nav-item mb-1">
              <a
                class="nav-link d-flex justify-content-between align-items-center"
                href="javascript:void(0)"
              >
                npm
                <span class="badge badge-pill badge-secondary ml-1">3篇</span>
              </a>
            </li>
            <li class="nav-item mb-1">
              <a
                class="nav-link d-flex justify-content-between align-items-center"
                href="javascript:void(0)"
              >
                npm
                <span class="badge badge-pill badge-secondary ml-1">3篇</span>
              </a>
            </li>
            <li class="nav-item mb-1">
              <a
                class="nav-link d-flex justify-content-between align-items-center"
                href="javascript:void(0)"
              >
                npm
                <span class="badge badge-pill badge-secondary ml-1">3篇</span>
              </a>
            </li>
          </ul>
          <div class="text-center push">
              <button type="button" class="btn btn-sm btn-light">查看所有标签</button>
          </div>
        </TBlock>
      </div>
    </div>
    <div class="col-md-7 col-xl-9">
      <TBlock title="全部文章">
        <template slot="options">
          <TButton icon="plus"></TButton>
          <TButton icon="magnifier" iconPack="si"></TButton>
          <TButton icon="arrow-left" iconPack="si"></TButton>
          <TButton icon="arrow-right" iconPack="si"></TButton>
          <TButton icon="refresh" iconPack="si"></TButton>
        </template>
        <div class="d-flex justify-content-between push">
          <span>
            <TButtonGroup size="sm" type="light" v-slot:default="{size, type, theme}" class="mr-1" >
              <TButton :size="size" :type="type" :theme="theme" icon="check-circle" iconPack="far" iconType="primary" v-if="!isSelectAll" @click="selectAll(true)">全选</TButton>
              <TButton :size="size" :type="type" :theme="theme" icon="circle" iconPack="far" iconType="primary" v-if="isSelectAll" @click="selectAll(false)">反选</TButton>
            </TButtonGroup>
            <TButton v-if="deleteIds.length > 0" size="sm" type="light" icon="times" iconType="danger">
              删除这{{deleteIds.length}}篇文章
            </TButton>
          </span>
          <ul class="pagination pagination-sm" style="margin-bottom: 0">
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
        </div>
        <div class="pull-x">
          <TTable :column="column" :data="data" selectKey="auth" hideHeader @onselected="selectedHandler" ref="dataTable" class="table table-hover table-vcenter font-size-sm">
            <template v-slot:auth="{item}" class="d-none d-sm-table-cell font-w600">
              {{item.auth}}
            </template>
            <template v-slot:title="{item, index}">
              <TBadge size="sm" type="success">未归类</TBadge>
              <a
                class="font-w600"
                data-toggle="modal"
                data-target="#one-inbox-message"
                href="#"
              > {{ item.title }}</a>
              <div class="text-muted mt-1">{{ index }} We are glad you decided to go with a vip..We are glad you decided to go with a vip..We are glad you decided to go with a vip..We are glad you decided to go with a vip..We are glad you decided to go with a vip..We are glad you decided to go with a vip..</div>
            </template>
            <template v-slot:tags="{item, index}">
              <i class="fa fa-tags mr-1"></i> {{ index }}, <a href="#">计算机</a>, NPM, JAVA
            </template>
            <template v-slot:time="{item, index}">
              <em>{{ index }} min ago</em>
            </template>
          </TTable>
        </div>
      </TBlock>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      column: [
        {
          title: '作者',
          index: 'auth',
          slot: 'auth',
          width: '140px',
          class: 'd-none d-sm-table-cell font-w600'
        },
        {
          title: '标题',
          index: 'title',
          slot: 'title'
        },
        {
          title: '标签',
          index: 'tags',
          slot: 'tags',
          width: '120px',
          class: 'd-none d-xl-table-cell text-muted'
        },
        {
          title: '时间',
          index: 'time',
          slot: 'time',
          width: '120px',
          class: 'd-none d-xl-table-cell text-muted'
        }
      ],
      data: [
        {
          auth: 'Justin Hunt',
          title: 'Welcome to our service',
          tags: '123',
          time: 'sdf'
        },
        {
          auth: 'Justin Hunt',
          title: 'Welcome to our service',
          tags: '123',
          time: 'sdf'
        },
        {
          auth: 'Justin Hunt',
          title: 'Welcome to our service',
          tags: '123',
          time: 'sdf'
        }
      ],
      isSelectAll: false,
      deleteIds: []
    }
  },
  methods: {
    showCategory() {
      const nav = this.$refs.category;
      const clsList = nav.classList;
      if (clsList.contains("d-none")) {
        clsList.remove("d-none");
      } else {
        clsList.add("d-none");
      }
    },
    selectAll(val) {
      this.$refs.dataTable.selectAll(val)
      this.isSelectAll = val
    },
    selectedHandler(vals, isAll) {
      this.deleteIds = vals
      this.isSelectAll = isAll
    }
  }
};
</script>