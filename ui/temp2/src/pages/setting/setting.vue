<template>
  <t-block title="系统设置" border theme>
    <b-form>
      <h2 class="content-heading pt-0">基础设置</h2>
      <b-row>
        <b-col :cols="4">
          <p class="text-muted">
              You can easily change the date format to match your preference
          </p>
        </b-col>
        <b-col :cols="4">
          <b-form-group
            description="Let us know your name."
            label="站点名称"
          >
            <b-form-input></b-form-input>
          </b-form-group>
          <b-form-group
            description="Let us know your name."
            label="域名"
          >
            <b-form-input></b-form-input>
          </b-form-group>
          <b-form-group
            description="Let us know your name."
            label="站点Logo"
          >
            <t-image-upload></t-image-upload>
          </b-form-group>
        </b-col>
      </b-row>
      <h2 class="content-heading pt-0">SEO设置</h2>
      <b-row>
        <b-col :cols="4">
          <p class="text-muted">
              You can easily change the date format to match your preference
          </p>
        </b-col>
        <b-col :cols="4">
          <b-form-group
            description="Let us know your name."
            label="SEO标题"
          >
            <b-form-input></b-form-input>
          </b-form-group>
          <b-form-group
            description="Let us know your name."
            label="SEO描述"
          >
            <b-form-textarea></b-form-textarea>
          </b-form-group>
          <b-form-group
            description="Let us know your name."
            label="SEO关键字"
          >
            <b-form-input></b-form-input>
          </b-form-group>
          <b-form-group
            description="Let us know your name."
            label="爬虫设置"
          >
            <b-form-textarea></b-form-textarea>
          </b-form-group>
        </b-col>
      </b-row>
      <h2 class="content-heading pt-0">系统参数设置</h2>
      <b-row>
        <b-col :cols="4">
          <p class="text-muted">
              You can easily change the date format to match your preference
          </p>
        </b-col>
        <b-col :cols="4">
          <b-form-group
            description="编辑和发布文章时的自动保存频率"
            label="自动保存频率"
          >
            <b-form-select>
              <b-form-select-option>30秒</b-form-select-option>
              <b-form-select-option>1分钟</b-form-select-option>
              <b-form-select-option>1分30秒</b-form-select-option>
              <b-form-select-option>2分钟</b-form-select-option>
            </b-form-select>
          </b-form-group>
          <b-form-group
            description="开启缓存可提高页面访问速度，减少数据库请求时间，但页面更新会有延迟"
            label="缓存开启/关闭"
          >
            <b-form-checkbox switch>
              关闭
            </b-form-checkbox>
          </b-form-group>
        </b-col>
      </b-row>
    </b-form>
  </t-block>
</template>
<script>
export default {
  data () {
    return {
      title: '系统设置',
      description: '可配置全局系统',
    }
  }
}
</script>