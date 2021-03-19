<template>
  <div class="createPost-container">
    <el-form ref="postForm" :model="postForm" :rules="rules" label-suffix=":" class="form-container">
      <sticky :z-index="10" :class-name="'sub-navbar '+postForm.status">
        <router-link :to="'/article/list'">
          <el-button style="float: left; margin-top: 7px; margin-left: 20px;">
            {{ $t('common.back') }}
          </el-button>
        </router-link>
        <el-button v-if="postForm.published != 1" v-loading="loading" style="margin-left: 10px;" type="success" @click="submitForm(1)">
          {{ $t('table.publish') }}
        </el-button>
        <el-button v-if="postForm.published != 1" v-loading="loading" type="warning" @click="draftForm">
          {{ $t('article.saveDraft') }}
        </el-button>
      </sticky>

      <div class="createPost-main-container">
        <el-row>
          <el-col :span="24">
            <el-form-item style="margin-bottom: 40px;" prop="title">
              <MDinput v-model="postForm.title" :maxlength="100" name="name" required>
                {{ $t('table.title') }}
              </MDinput>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <div class="postInfo-container">
            <el-row>
              <el-col :span="12">
                <el-form-item prop="tags" label-width="60px" :label="$t('article.tag')" class="postInfo-container-item">
                  <el-select v-model="postForm.tags" multiple :placeholder="$t('common.pleaseChoose')" style="width: 300px;">
                    <el-option v-for="(item) in tagListOptions" :key="item.id" :label="item.name" :value="item.id" />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label-width="60px" label="SEO" class="postInfo-container-item" style="width: 400px;">
                  <el-input v-model="postForm.keywords" />
                </el-form-item>
              </el-col>
            </el-row>
          </div>
        </el-row>

        <el-form-item prop="description" style="margin-bottom: 40px;" label-width="60px" :label="$t('article.description')">
          <el-input v-model="postForm.description" :rows="1" type="textarea" class="article-textarea" autosize :placeholder="$t('common.pleaseEnter')" />
          <span v-show="descriptionLength" class="word-counter">{{ descriptionLength }}words</span>
        </el-form-item>

        <el-form-item prop="content" style="margin-bottom: 30px;">
          <markdown-editor ref="markdownEditor" v-model="postForm.content" :language="lang" :options="{}" height="500px" />
        </el-form-item>

        <!-- <el-form-item prop="image_uri" style="margin-bottom: 30px;">
          <Upload v-model="postForm.image_uri" />
        </el-form-item> -->
      </div>
    </el-form>
  </div>
</template>

<script>
import MarkdownEditor from '@/components/MarkdownEditor'
// import Upload from '@/components/Upload/SingleImage3'
import MDinput from '@/components/MDinput'
import Sticky from '@/components/Sticky' // 粘性header组件
import { fetchArticle, createArticle, updateArticle } from '@/api/article'
import { fetchSelectList } from '@/api/tag'

const defaultForm = {
  title: '', // 文章题目
  cover: '', // 文章封面
  keywords: '', // 文章关键词
  description: '', // 文章摘要
  content: '', // 文章内容
  type: 1, // 文章类型
  tags: [], // 文章管理标签
  id: undefined
}

export default {
  name: 'ArticleDetail',
  components: {
    MarkdownEditor,
    MDinput,
    // Upload,
    Sticky
  },
  props: {
    isEdit: {
      type: Boolean,
      default: false
    }
  },
  data() {
    // const validateRequire = (rule, value, callback) => {
    //   if (value === '') {
    //     this.$message({
    //       message: rule.field + '为必传项',
    //       type: 'error'
    //     })
    //     callback(new Error(rule.field + '为必传项'))
    //   } else {
    //     callback()
    //   }
    // }
    return {
      postForm: Object.assign({}, defaultForm),
      loading: false,
      tagListOptions: [],
      rules: {
        title: [{ required: true, trigger: 'blur' }],
        content: [{ required: true, trigger: 'blur' }],
        tags: [{ required: true, trigger: 'change' }],
        description: [{ required: true, trigger: 'change' }]
        // source_uri: [{ validator: validateSourceUri, trigger: 'blur' }]
      },
      tempRoute: {},
      html: ''
    }
  },
  computed: {
    descriptionLength() {
      return this.postForm.description.length
    },
    lang() {
      return this.$store.getters.language
    }
  },
  created() {
    this.getTagList()
    if (this.isEdit) {
      const id = this.$route.params && this.$route.params.id
      this.fetchData(id)
    }
    this.tempRoute = Object.assign({}, this.$route)
  },
  methods: {
    fetchData(id) {
      fetchArticle(id).then(response => {
        const tags = []
        if (response.data.tags) {
          response.data.tags.forEach(item => {
            tags.push(item.id)
          })
        }
        response.data.tags = tags
        this.postForm = response.data
        // set tagsview title
        this.setTagsViewTitle()
        // set page title
        this.setPageTitle()
      }).catch(err => {
        console.log(err)
      })
    },
    setTagsViewTitle() {
      const title = this.$t('route.editArticle')
      const route = Object.assign({}, this.tempRoute, { title: `${title}-${this.postForm.id}` })
      this.$store.dispatch('tagsView/updateVisitedView', route)
    },
    setPageTitle() {
      const title = this.$t('route.editArticle')
      document.title = `${title} - ${this.postForm.id}`
    },
    submitForm(published) {
      published = parseInt(published)
      // this.html = this.$refs.markdownEditor.getHtml()
      this.$refs.postForm.validate(valid => {
        if (valid) {
          const tempData = Object.assign({}, this.postForm)
          tempData.tags = tempData.tags.join(',')
          tempData.published = published
          if (tempData.id) {
            this.updateData(tempData)
          } else {
            this.createData(tempData)
          }
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    draftForm() {
      this.submitForm(0)
    },
    createData(data) {
      this.loading = true
      createArticle(data).then(() => {
        this.loading = false
        this.$message({
          message: this.$t(data.published ? 'common.publishSuccess' : 'common.saveSuccess'),
          type: 'success',
          showClose: true,
          duration: 1000
        })
        if (data.published) {
          this.$router.push('/article/list')
        } else {
          // TODO 传参跳转至编辑页面

          this.$router.push('/article/edit', {})
        }
      }).catch(() => {
        this.loading = false
        this.$message({
          message: this.$t('common.operationFailed'),
          type: 'error',
          showClose: true,
          duration: 2000
        })
      })
    },
    updateData(data) {
      this.loading = true
      updateArticle(data.id, data).then(() => {
        this.loading = false
        this.$message({
          message: this.$t(data.published ? 'common.publishSuccess' : 'common.saveSuccess'),
          type: 'success',
          showClose: true,
          duration: 1000
        })
        if (data.published) {
          this.$router.push('/article/list')
        }
      }).catch(() => {
        this.loading = false
        this.$message({
          message: this.$t('common.operationFailed'),
          type: 'error',
          showClose: true,
          duration: 2000
        })
      })
    },
    getTagList(query) {
      query = query || {}
      fetchSelectList(query).then(response => {
        if (!response.data) return
        this.tagListOptions = response.data.map(v => {
          return { id: v.id, name: v.name }
        })
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import "~@/styles/mixin.scss";

.createPost-container {
  position: relative;

  .createPost-main-container {
    padding: 40px 45px 20px 50px;

    .postInfo-container {
      position: relative;
      @include clearfix;
      margin-bottom: 10px;

      .postInfo-container-item {
        float: left;
      }
    }
  }

  .word-counter {
    width: 40px;
    position: absolute;
    right: 10px;
    top: 0px;
  }
}

.article-textarea ::v-deep {
  textarea {
    padding-right: 40px;
    resize: none;
    border: none;
    border-radius: 0px;
    border-bottom: 1px solid #bfcbd9;
  }
}
</style>
