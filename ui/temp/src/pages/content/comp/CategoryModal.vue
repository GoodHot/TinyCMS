<template>
  <b-form>
    <b-form-group label="上级分类" >
      <b-dropdown :text="categoryName" variant="outline-secondary" class="t-create-category" style="width: 100%" >
        <TCategory :defaultCategory="defaultCategory" @selectItem="selectItem" :selected="categorySelected"></TCategory>
      </b-dropdown>
    </b-form-group>

    <b-form-group label="分类名称" >
      <b-form-input placeholder="分类名称" :state="form.name.valid" v-model="form.name.value"></b-form-input>
      <b-form-invalid-feedback :state="form.name.valid">
        {{form.name.msg}}
      </b-form-invalid-feedback>
    </b-form-group>

    <b-form-group label="分类路径">
      <b-form-input placeholder="分类路径" :state="form.path.valid" v-model="form.path.value"></b-form-input>
      <b-form-invalid-feedback :state="form.path.valid">
        {{form.path.msg}}
      </b-form-invalid-feedback>
    </b-form-group>

    <b-form-group label="排序" description="数字越大越靠前">
      <b-form-input placeholder="排序" :state="form.sort.valid" v-model="form.sort.value"></b-form-input>
      <b-form-invalid-feedback :state="form.sort.valid">
        {{form.sort.msg}}
      </b-form-invalid-feedback>
    </b-form-group>

    <b-form-group label="SEO标题">
      <b-form-input placeholder="SEO标题" v-model="form.seo_title.value"></b-form-input>
      <b-form-invalid-feedback :state="form.seo_title.valid">
        {{form.seo_title.msg}}
      </b-form-invalid-feedback>
    </b-form-group>

    <b-form-group label="SEO关键字">
      <b-form-input placeholder="SEO关键字" v-model="form.keyword.value"></b-form-input>
      <b-form-invalid-feedback :state="form.keyword.valid">
        {{form.keyword.msg}}
      </b-form-invalid-feedback>
    </b-form-group>

    <b-form-group label="SEO描述">
      <b-form-textarea placeholder="SEO描述" rows="3" max-rows="6" v-model="form.description.value"></b-form-textarea>
      <b-form-invalid-feedback :state="form.description.valid">
        {{form.description.msg}}
      </b-form-invalid-feedback>
    </b-form-group>

    <b-form-group label="备注">
      <b-form-textarea placeholder="备注" rows="3" max-rows="6" v-model="form.remark.value"></b-form-textarea>
      <b-form-invalid-feedback :state="form.remark.valid">
        {{form.remark.msg}}
      </b-form-invalid-feedback>
    </b-form-group>

    <b-form-group label="分类图标">
      <TUpload></TUpload>
    </b-form-group>
  </b-form>
</template>
<script>
import valid from '@/utils/valid'
import {saveCategory,getCategory} from '@/api/category'

export default {
  props: {
    editId: {
      type: Number,
      default: 0
    }
  },
  data() {
    return {
      form: {
        id: {
          value: 0
        },
        parent_id: {
          value: null
        },
        name: {
          value: '',
          require: '请输入分类名称'
        },
        path: {
          value: '',
          require: '请输入分类路径',
          regex: {
            reg: "^\\w+$",
            msg: '只能输入字母和数字'
          }
        },
        seo_title: {
          value: ''
        },
        description: {
          value: ''
        },
        keyword: {
          value: ''
        },
        remark: {
          value: ''
        },
        sort: {
          value: 0
        },
        icon: {
          value: ''
        }
      },
      defaultCategory: [
        {
          id: null,
          title: "顶级分类",
          icon: "list-ul",
          remark: "0篇",
          active: false
        }
      ],
      categoryName: '顶级分类',
      categorySelected: null
    }
  },
  mounted() {
    this.loadCategory()
  },
  methods: {
    submit() {
      const rst = valid.check(this.form)
      this.form = rst.values
      if (!rst.valid) {
        return false
      }
      const param = {
        id: this.form.id.value,
        name: this.form.name.value,
        path: this.form.path.value,
        seo_title: this.form.seo_title.value,
        description: this.form.description.value,
        keyword: this.form.keyword.value,
        remark: this.form.remark.value,
        icon: this.form.icon.value,
        parent_id: this.form.parent_id.value,
        sort: parseInt(this.form.sort.value)
      }
      saveCategory(param).then(() => {
        this.$emit('savedCategory')
      })
    },
    selectItem(row) {
      this.form.parent_id.value = row.id
      this.categoryName = row.title
    },
    loadCategory() {
      if (this.editId == 0) {
        return
      }
      getCategory(this.editId).then(res => {
        const cat = res.category
        this.form.id.value = cat.id
        this.form.parent_id.value = cat.parent_id
        this.form.name.value = cat.name
        this.form.path.value = cat.path
        this.form.seo_title.value = cat.seo_title
        this.form.description.value = cat.description
        this.form.keyword.value = cat.keyword
        this.form.remark.value = cat.remark
        this.form.sort.value = cat.sort
        this.form.icon.value = cat.icon
        this.categorySelected = cat.parent_id
      })
    }
  }
}
</script>