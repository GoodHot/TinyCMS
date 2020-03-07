<template>
  <TBlock title="分类管理">
    <div class="pull-x">
      <TTable :column="column" :data="data" ref="dataTable" class="table table-hover table-vcenter">
        <template v-slot:name="{item}">
          <span v-for="i in item.indent" :key="i" style="margin-left: 20px"></span>
          {{item.name}}
        </template>
        <template v-slot:action="{item}">
          <b-button-group size="sm">
            <b-button variant="light" v-b-tooltip.hover title="编辑" @click="onedit(item.id)"><TIcon icon="settings" pack="si"></TIcon></b-button>
            <b-button variant="light" v-b-tooltip.hover title="删除"><TIcon icon="trash" pack="si"></TIcon></b-button>
          </b-button-group>
        </template>
      </TTable>
    </div>
    <b-modal
      v-model="categoryVisible"
      size="lg"
      title="编辑分类"
      no-close-on-backdrop
      @ok="submitCategory"
      ok-title="保存"
      cancel-title="关闭"
    >
      <CategoryModal ref="categoryModal" @savedCategory="savedCategoryHandler" :editId="editId"></CategoryModal>
    </b-modal>
  </TBlock>
</template>
<script>
import {categoryTree} from '@/api/category'
import CategoryModal from "./comp/CategoryModal";

export default {
  components: {
    CategoryModal
  },
  data() {
    return {
      editId: 0,
      categoryVisible: false,
      column: [
        {
          title: 'ID',
          index: 'id',
          align: 'center',
          width: '80px'
        },
        {
          title: '分类名称',
          index: 'name',
          width: '300px',
          slot: 'name'
        },
        {
          title: '路径',
          index: 'path',
          width: '150px'
        },
        {
          title: '排序',
          index: 'sort',
          width: '150px'
        },
        {
          title: '关键字',
          index: 'keyword',
          width: '350px'
        },
        {
          title: '备注',
          index: 'remark'
        },
        {
          title: '文章数',
          index: 'article_count',
          width: '120px'
        },
        {
          title: '操作',
          slot: 'action',
          width: '150px'
        }
      ],
      data: []
    }
  },
  mounted() {
    this.loadCategoryTree()
  },
  methods: {
    iterationCategory(trees, deep) {
      for (let i in trees) {
        const tree = trees[i]
        tree.channel.indent = deep
        this.data.push(tree.channel)
        if (tree.children) {
          this.iterationCategory(tree.children, deep + 1)
        }
      }
    },
    loadCategoryTree() {
      this.data = []
      categoryTree().then(res => {
        this.iterationCategory(res.tree, 0)
      })
    },
    submitCategory(bvModalEvt) {
      bvModalEvt.preventDefault();
      this.$refs.categoryModal.submit();
      return false;
    },
    onedit(id) {
      this.categoryVisible = true
      this.editId = id
    },
    savedCategoryHandler() {
      this.$bvToast.toast('袖肥分类成功', {
        title: '提示',
        variant: 'info',
        solid: true
      })
      this.categoryVisible = false
      this.loadCategoryTree()
    }
  }
}
</script>