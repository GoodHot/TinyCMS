<template>
  <div class="row">
    <div class="col-md-5 col-xl-3">
      <div class="d-md-none push">
        <b-button variant="primary" class="btn btn-block btn-primary" @click="showCategory">文章分类</b-button>
      </div>
      <div ref="category" class="d-none d-md-block push">
        <TBlock title="文章分类" headerBg>
          <template slot="options">
            <b-button-group size="sm">
              <b-button class="btn-block-option" variant="light" @click="categoryVisible = !categoryVisible"><TIcon icon="plus" /></b-button>
              <b-button class="btn-block-option" variant="light"><TIcon icon="settings" pack="si" /></b-button>
            </b-button-group>
          </template>
          <TList class="push" :data="category"></TList>
        </TBlock>
        <TBlock title="标签" icon="tags">
          <TList class="push" :data="category"></TList>
          <div class="text-center push">
            <b-button size="sm" variant="light">查看所有标签</b-button>
          </div>
        </TBlock>
      </div>
    </div>
    <div class="col-md-7 col-xl-9">
      <TBlock title="全部文章">
        <template slot="options">
          <b-button-group size="sm">
            <b-button class="btn-block-option" variant="light" @click="$router.push('/publish')"><TIcon icon="plus" /></b-button>
            <b-button class="btn-block-option" variant="light"><TIcon icon="magnifier" pack="si" /></b-button>
            <b-button class="btn-block-option" variant="light"><TIcon icon="arrow-left" pack="si" /></b-button>
            <b-button class="btn-block-option" variant="light"><TIcon icon="arrow-right" pack="si" /></b-button>
            <b-button class="btn-block-option" variant="light"><TIcon icon="refresh" pack="si" /></b-button>
          </b-button-group>
        </template>
        <div class="d-flex justify-content-between push">
          <span>
            <b-button-group size="sm" class="mr-1">
              <b-button variant="light"><TIcon class="text-primary" icon="check-circle" pack="far" /> 全选</b-button>
              <b-button variant="light"><TIcon class="text-primary" icon="circle" pack="far" /> 反选</b-button>
            </b-button-group>
            <b-button variant="light" size="sm"><TIcon class="text-danger" icon="times"  /> 删除这{{deleteIds.length}}篇文章</b-button>
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
          <ArticleTable :pagination="true" :column="column" :data="data" selectKey="auth" hideHeader @onselected="selectedHandler" ref="dataTable" class="table table-hover table-vcenter font-size-sm">
            <template v-slot:auth="{item}" class="d-none d-sm-table-cell font-w600">
              {{item.auth}}
            </template>
            <template v-slot:title="{item, index}">
              <b-badge variant="success">未归类</b-badge>
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
          </ArticleTable>
        </div>
      </TBlock>
    </div>
    <b-modal v-model="categoryVisible" size="lg" title="创建分类" no-close-on-backdrop>
      <CategoryModal></CategoryModal>
    </b-modal>
  </div>
</template>
<script>
import ArticleTable from './comp/ArticleTable'
import CategoryModal from './comp/CategoryModal'
export default {
  components: {
    ArticleTable,
    CategoryModal
  },
  data() {
    return {
      categoryVisible: false,
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
      deleteIds: [],
      category: [
        {
          title: '全部分类',
          icon: 'book',
          remark: '3篇',
          img: 'http://ww1.sinaimg.cn/mw1024/00745YaMgy1gbyi3gl2mhj33402c0hdw.jpg'
        },
        {
          title: '全部分类1',
          icon: 'book',
          remark: '3篇'
        },
        {
          title: '全部分类2',
          icon: 'book',
          remark: '3篇'
        },
        {
          title: '全部分类3',
          icon: 'book',
          remark: '3篇',
          children: [
            {
              title: '全部分类4',
              icon: 'book',
              remark: '3篇',
              children: [
                {
                  title: '全部分类4',
                  icon: 'book',
                  remark: '3篇'
                },
                {
                  title: '全部分类5',
                  icon: 'book',
                  remark: '3篇'
                }
              ]
            },
            {
              title: '全部分类5',
              icon: 'book',
              remark: '3篇'
            }
          ]
        }
      ]
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