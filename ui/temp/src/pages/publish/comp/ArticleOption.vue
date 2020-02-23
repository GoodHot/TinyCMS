<template>
  <div class="t-publish-option">
    <div class="options">
      <b-dropdown :text="categoryName" variant="light" class="mr-2 float-left t-publish-option-category">
        <TCategory ref="categoryList" @selectItem="selectItem"></TCategory>
      </b-dropdown>
      <b-button class="mr-2" variant="light" @click="$emit('showAdvance', $event)">
        <TIcon icon="settings" pack="si" />
        高级设置
      </b-button>
      <TagInput ref="tags"></TagInput>
    </div>
    <div>
      <b-button class="mr-2" variant="success" @click="$emit('onpublish', $event)">
        <TIcon icon="paper-plane" pack="far" />
        立即发布
      </b-button>

      <b-button variant="light" @click="$emit('onsave', $event)">
        <TIcon icon="save" pack="far" />
        保存为草稿
      </b-button>
    </div>
  </div>
</template>
<script>
import TagInput from './TagInput'
export default {
  components:{
    TagInput
  },
  data() {
    return {
      categoryName: '选择分类',
      categoryId: null
    }
  },
  methods: {
    selectItem(row) {
      this.categoryName = ` 分类: ${row.title} `
      this.categoryId = row.id
    },
    getValue() {
      const tags = this.$refs.tags.getValue()
      return {
        category_id: this.categoryId,
        tags
      }
    }
  }
}
</script>