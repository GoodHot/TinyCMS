<template>
  <div>
    <h1>elements</h1>
    <h2>theme</h2>
    <TInput placeholder="default" />
    <br>
    <TInput placeholder="alt" theme="alt" />
    <h2>readonly</h2>
    <TInput placeholder="alt" readonly />
    <br>
    <TInput placeholder="alt" readonly theme="alt" />
    <h2>disabled</h2>
    <TInput placeholder="alt" disabled />
    <br>
    <TInput placeholder="alt" disabled theme="alt" />
    <h2>state</h2>
    <TInput placeholder="success" state="success" />
    <br>
    <TInput placeholder="success" state="error" />
    <br>
    <TInput placeholder="success" state="success" theme="alt" />
    <br>
    <TInput placeholder="success" state="error" theme="alt" />
    <h2>textarea</h2>
    <TTextarea>hello</TTextarea><br>
    <TTextarea theme="alt">hello</TTextarea><br>
    <TTextarea state="success">hello</TTextarea><br>
    <TTextarea state="error">hello</TTextarea><br>
    <TTextarea theme="alt" state="success">hello</TTextarea><br>
    <TTextarea theme="alt" state="error">hello</TTextarea>
    <h2>switch</h2>
    <TSwitch v-for="item of types" :key="item" :type="item" :name="item" :label="item" /><br>
    <TSwitch v-for="item of types" :key="item + '1'" :type="item" :name="item + '1'" :label="item" size="lg" /><br>
    <TSwitch v-for="item of types" :key="item + '2'" :type="item" :name="item + '2'" :label="item" size="sm" /><br>
    <TSwitch label="aaa" :inline="true" /><TSwitch label="bbb" :inline="true" /><TSwitch label="ccc" :inline="true" />
    <h2>checkbox</h2>
    <TCheckbox v-for="item of types" :key="item + '3'" :type="item" :name="item" :label="item" /><br>
    <TCheckbox v-for="item of types" :key="item + '4'" :type="item" :name="item + '3'" :label="item" size="lg" shape="square" /><br>
    <TCheckbox v-for="item of types" :key="item + '5'" :type="item" :name="item + '5'" :label="item" size="sm" /><br>
    <TCheckbox label="aaa" :inline="true" checked /><TCheckbox label="bbb" :inline="true" /><TCheckbox label="ccc" :inline="true" checked /><br>
    <TCheckbox label="aaa" :inline="true" checked shape="square" /><TCheckbox label="bbb" :inline="true"  shape="square" /><TCheckbox label="ccc" :inline="true" checked shape="square" />
    <h1>form</h1>
    <div class="row">
      <div class="col-6">
        <TBlock title="form" headerBg theme="primary">
          <TForm>
            <TFormItem label="Email address" help="We'll never share your email with anyone else." v-slot:default="item" >
              <TInput placeholder="请输入邮箱" :state="item.state" />
            </TFormItem>
          </TForm>
        </TBlock>
      </div>
      <div class="col-6">
        <TBlock title="form" headerBg theme="primary">
          <p>sdfasdfsdfs</p>
        </TBlock>
      </div>
      <div class="col-4">
        <TBlock title="form" headerBg theme="primary">
          <p>sdfasdfsdfs</p>
        </TBlock>
      </div>
      <div class="col-4">
        <TBlock title="form" headerBg theme="primary">
          <p>sdfasdfsdfs</p>
        </TBlock>
      </div>
      <div class="col-4">
        <TBlock title="form" headerBg theme="primary">
          <p>sdfasdfsdfs</p>
        </TBlock>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      types: ['primary','secondary','success','info','warning','danger','dark','light']
    }
  }
}
</script>