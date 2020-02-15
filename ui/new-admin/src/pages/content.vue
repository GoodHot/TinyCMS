<template>
  <div class="row">
    <div class="col-md-5 col-xl-3">
      <div class="d-md-none push">
        <button type="button" class="btn btn-block btn-primary" @click="showCategory">文章分类</button>
      </div>
      <div ref="category" class="d-none d-md-block push">
        <TBlock title="文章分类" headerBg>
          <template slot="options">
            <button type="button" class="btn btn-success btn-sm">
              <i class="fa fa-plus ml-1 fa-fw"></i>新建分类
            </button>
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
          </ul>
        </TBlock>
      </div>
    </div>
    <div class="col-md-7 col-xl-9">
      <TBlock title="全部文章">
        <template slot="options">
          <TButton icon="magnifier" iconPack="si"></TButton>
          <TButton icon="arrow-left" iconPack="si"></TButton>
          <TButton icon="arrow-right" iconPack="si"></TButton>
          <TButton icon="refresh" iconPack="si"></TButton>
        </template>
        <div class="d-flex justify-content-between push">
          <TButtonGroup size="sm" type="light">
            <TButton icon="archive" iconType="primary">Archive</TButton>
            <TButton icon="star" iconType="warning">Archive</TButton>
          </TButtonGroup>
          <button class="btn btn-sm btn-light" type="button">
            <i class="fa fa-times text-danger"></i>
            <span class="d-none d-sm-inline ml-1">Delete</span>
          </button>
        </div>
        <div class="pull-x">
          <table
            class="js-table-checkable table table-hover table-vcenter font-size-sm js-table-checkable-enabled"
          >
            <tbody>
              <tr class>
                <td class="text-center" style="width: 60px;">
                  <div class="custom-control custom-checkbox">
                    <input
                      type="checkbox"
                      class="custom-control-input"
                      id="inbox-msg15"
                      name="inbox-msg15"
                    />
                    <label class="custom-control-label font-w400" for="inbox-msg15"></label>
                  </div>
                </td>
                <td class="d-none d-sm-table-cell font-w600" style="width: 140px;">Justin Hunt</td>
                <td>
                  <a
                    class="font-w600"
                    data-toggle="modal"
                    data-target="#one-inbox-message"
                    href="#"
                  >Welcome to our service</a>
                  <div class="text-muted mt-1">It's a pleasure to have you on our service..</div>
                </td>
                <td class="d-none d-xl-table-cell text-muted" style="width: 80px;">
                  <i class="fa fa-paperclip mr-1"></i> (3)
                </td>
                <td class="d-none d-xl-table-cell text-muted" style="width: 120px;">
                  <em>2 min ago</em>
                </td>
              </tr>
              <tr class>
                <td class="text-center">
                  <div class="custom-control custom-checkbox">
                    <input
                      type="checkbox"
                      class="custom-control-input"
                      id="inbox-msg14"
                      name="inbox-msg14"
                    />
                    <label class="custom-control-label font-w400" for="inbox-msg14"></label>
                  </div>
                </td>
                <td class="d-none d-sm-table-cell font-w600">Jose Mills</td>
                <td>
                  <a
                    class="font-w600"
                    data-toggle="modal"
                    data-target="#one-inbox-message"
                    href="#"
                  >Your subscription was updated</a>
                  <div class="text-muted mt-1">We are glad you decided to go with a vip..</div>
                </td>
                <td class="d-none d-xl-table-cell text-muted">
                  <i class="fa fa-paperclip mr-1"></i> (2)
                </td>
                <td class="d-none d-xl-table-cell text-muted">
                  <em>10 min ago</em>
                </td>
              </tr>
              <tr class>
                <td class="text-center">
                  <div class="custom-control custom-checkbox">
                    <input
                      type="checkbox"
                      class="custom-control-input"
                      id="inbox-msg13"
                      name="inbox-msg13"
                    />
                    <label class="custom-control-label font-w400" for="inbox-msg13"></label>
                  </div>
                </td>
                <td class="d-none d-sm-table-cell font-w600">Jeffrey Shaw</td>
                <td>
                  <a
                    class="font-w600"
                    data-toggle="modal"
                    data-target="#one-inbox-message"
                    href="#"
                  >Update is available</a>
                  <div class="text-muted mt-1">An update is under way for your app..</div>
                </td>
                <td class="d-none d-xl-table-cell text-muted"></td>
                <td class="d-none d-xl-table-cell text-muted">
                  <em>25 min ago</em>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </TBlock>
      
    </div>
  </div>
</template>
<script>
export default {
  methods: {
    showCategory() {
      const nav = this.$refs.category;
      const clsList = nav.classList;
      if (clsList.contains("d-none")) {
        clsList.remove("d-none");
      } else {
        clsList.add("d-none");
      }
    }
  }
};
</script>