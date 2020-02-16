<template>
  <TBlock title="table" subtitle="example" headerBg>
    <template slot="options">
      <TButton iconPack="fa" icon="plus" />
      <TButton iconPack="si" icon="magnifier" />
      <TButton iconPack="si" icon="refresh" />
    </template>
    <TForm>
      <TFormItem col="4" label="First name" v-slot:default="item" >
        <TInput placeholder="请输入邮箱" theme="alt" :state="item.state" />
      </TFormItem>
      <TFormItem col="4" label="First name" v-slot:default="item" >
        <TInput placeholder="请输入邮箱" theme="alt" :state="item.state" />
      </TFormItem>
      <TFormItem col="4" label="First name" v-slot:default="item" >
        <TInput placeholder="请输入邮箱" theme="alt" :state="item.state" />
      </TFormItem>
      <TFormItem col="4" label="First name" v-slot:default="item" >
        <TInput placeholder="请输入邮箱" theme="alt" :state="item.state" />
      </TFormItem>
      <TFormItem col="4" v-slot:default >
        <button type="submit" class="btn btn-alt-success">Update</button>
      </TFormItem>
    </TForm>
    <TTable :column="column" :data="data" selectKey="id" @onselected="selectedHandler">
      <template slot="action" scope>
        <a
          class="btn btn-sm btn-alt-secondary js-tooltip-enabled"
          href="be_pages_ecom_product_edit.html"
          data-toggle="tooltip"
          title
          data-original-title="View"
        >
          <i class="fa fa-fw fa-eye"></i>
        </a>
        <a
          class="btn btn-sm btn-alt-danger js-tooltip-enabled"
          href="javascript:void(0)"
          data-toggle="tooltip"
          title
          data-original-title="Delete"
        >
          <i class="fa fa-fw fa-times text-danger"></i>
        </a>
      </template>
      <template v-slot:state="{item, index}">
        <TBadge size="sm">{{item.status}} {{ index }}</TBadge>
      </template>
    </TTable>
  </TBlock>
</template>
<script>
export default {
  data() {
    return {
      // 'id', 'submitted', 'status', 'customer', 'products', 'value', 'action'
      column: [
        {
          title: "ID",
          index: "id",
          width: "10%",
          align: "center"
        },
        {
          title: "submitted",
          index: "submitted",
          width: "20%"
        },
        {
          title: "status",
          index: "status",
          width: "10%",
          slot: "state"
        },
        {
          title: "customer",
          index: "customer",
          width: "10%"
        },
        {
          title: "products",
          index: "products",
          width: "10%"
        },
        {
          title: "value",
          index: "value",
          width: "10%"
        },
        {
          title: "action",
          index: "action",
          align: "center",
          slot: "action"
        }
      ],
      data: [
        {
          id: "123",
          submitted: "sss",
          status: "s1",
          customer: "fff",
          products: "ggg",
          value: "666"
        },
        {
          id: "124",
          submitted: "sss",
          status: "s2",
          customer: "fff",
          products: "ggg",
          value: "666"
        },
        {
          id: "125",
          submitted: "sss",
          status: "s3",
          customer: "fff",
          products: "ggg",
          value: "666"
        },
        {
          id: "126",
          submitted: "sss",
          status: "s4",
          customer: "fff",
          products: "ggg",
          value: "666"
        }
      ]
    };
  },
  methods: {
    selectedHandler(vals) {
      console.log(vals)
    }
  }
};
</script>