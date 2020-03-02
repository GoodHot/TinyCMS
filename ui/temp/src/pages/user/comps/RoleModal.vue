<template>
  <b-form>
    <b-form-group label="角色编码" >
      <b-form-input placeholder="角色编码" :state="form.code.valid" v-model="form.code.value"></b-form-input>
      <b-form-invalid-feedback :state="form.code.valid">
        {{form.code.msg}}
      </b-form-invalid-feedback>
    </b-form-group>
    <b-form-group label="角色名称" >
      <b-form-input placeholder="角色名称" :state="form.role_name.valid" v-model="form.role_name.value"></b-form-input>
      <b-form-invalid-feedback :state="form.role_name.valid">
        {{form.role_name.msg}}
      </b-form-invalid-feedback>
    </b-form-group>
    <b-form-group label="角色名称" >
      <TTree :data="roleTree" expand></TTree>
    </b-form-group>
  </b-form>
</template>
<script>
export default {
  data() {
    return {
      form: {
        code: {
          value: '',
          require: '请输入角色编码'
        },
        role_name: {
          value: '',
          require: '请输入角色名称'
        }
      },
      roleTree: [
        {
          name: 'content',
          value: 111,
        },
        {
          name: 'content1',
          value: 222,
          checked: true
        },
        {
          name: 'content2',
          value: 333,
          child: [
            {
              name: 'content',
              value: 111,
            },
            {
              name: 'content1',
              value: 222,
              checked: true,
              child: [
                {
                  name: 'content',
                  value: 111,
                },
                {
                  name: 'content1',
                  value: 222,
                  checked: true
                }
              ]
            }
          ]
        },
        {
          name: 'content3',
          value: 444,
        },
      ]
    }
  }
}
</script>