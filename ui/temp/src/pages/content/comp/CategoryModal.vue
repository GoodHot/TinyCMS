<template>
  <b-form>
    <b-form-group label="上级分类" >
      <b-form-select v-model="form.parent_id.value" :options="categoryTree"></b-form-select>
      <b-form-invalid-feedback :state="form.parent_id.valid">
        {{form.parent_id.msg}}
      </b-form-invalid-feedback>
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
      <b-form-textarea placeholder="SEO描述" rows="3" max-rows="6" v-model="form.desccription.value"></b-form-textarea>
      <b-form-invalid-feedback :state="form.desccription.valid">
        {{form.desccription.msg}}
      </b-form-invalid-feedback>
    </b-form-group>

    <b-form-group label="备注">
      <b-form-textarea placeholder="备注" rows="3" max-rows="6" v-model="form.remark.value"></b-form-textarea>
      <b-form-invalid-feedback :state="form.remark.valid">
        {{form.remark.msg}}
      </b-form-invalid-feedback>
    </b-form-group>

    <b-form-group label="分类图标">
      <b-form-file
        placeholder="上传图片"
        drop-placeholder="Drop file here..."
      ></b-form-file>
      <b-form-invalid-feedback :state="form.icon.valid">
        {{form.icon.msg}}
      </b-form-invalid-feedback>
    </b-form-group>
  </b-form>
</template>
<script>
import valid from '@/utils/valid'
import {categoryTree, saveCategory} from '@/api/category'

export default {
  data() {
    return {
      form: {
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
        desccription: {
          value: ''
        },
        keyword: {
          value: ''
        },
        remark: {
          value: ''
        },
        icon: {
          value: ''
        }
      },
      categoryTree: [
        { value: null, text: '顶级分类' }
      ]
    }
  },
  mounted() {
    this.getCategoryTree()
  },
  methods: {
    submit() {
      const rst = valid.check(this.form)
      this.form = rst.values
      if (!rst.valid) {
        return false
      }
      const param = {
        name: this.form.name.value,
        path: this.form.path.value,
        seo_title: this.form.seo_title.value,
        desccription: this.form.desccription.value,
        keyword: this.form.keyword.value,
        remark: this.form.remark.value,
        icon: this.form.icon.value,
        parent_id: this.form.parent_id.value
      }
      saveCategory(param).then(() => {
        this.$emit('savedCategory')
      })
    },
    getCategoryTree() {
      categoryTree().then(res => {
        if (res.tree) {
          console.log(res.tree)
          this.setCategoryTree(0, res.tree)
        }
      })
    },
    setCategoryTree(level, tree) {
      let placeholder = ''
      for (let i = 0 ; i < level ; i++) {
        if (i === 0) {
          placeholder += '|--'
        } else {
          placeholder = '&nbsp;&nbsp;&nbsp;&nbsp;' + placeholder
        }
      }
      for (let i in tree) {
        const data = tree[i]
        this.categoryTree.push({
          value: data.id,
          html: placeholder + data.name
        })
        if (data.children && data.children.length > 0) {
          this.setCategoryTree((level + 1), data.children)
        }
      }
    }
  }
}
</script>