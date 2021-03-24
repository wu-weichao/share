<template>
  <div class="app-container">
    <div class="filter-container">
      <el-select v-model="listQuery.tags" multiple :placeholder="$t('common.pleaseChoose')" style="width: 300px;margin-right: 10px;" class="filter-item">
        <el-option v-for="(item) in tagListOptions" :key="item.id" :label="item.name" :value="item.id" />
      </el-select>
      <el-button v-waves class="filter-item" icon="el-icon-search" @click="handleFilter">
        {{ $t('table.search') }}
      </el-button>
      <router-link :to="'/article/create'">
        <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit">
          {{ $t('table.add') }}
        </el-button>
      </router-link>
    </div>

    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="ID" width="80">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column min-width="300px" :label="$t('table.title')">
        <template slot-scope="scope">
          <router-link :to="'/article/edit/'+scope.row.id" class="link-type">
            <span>{{ scope.row.title }}</span>
          </router-link>
        </template>
      </el-table-column>

      <el-table-column width="140px" align="center" :label="$t('common.createdAt')">
        <template slot-scope="scope">
          <span>{{ scope.row.created_at | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column width="140px" align="center" :label="$t('common.updatedAt')">
        <template slot-scope="scope">
          <span>{{ scope.row.updated_at | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column width="120px" align="center" :label="$t('article.tag')">
        <template slot-scope="scope">
          <span>{{ getTagNames(scope.row.tags) }}</span>
        </template>
      </el-table-column>

      <!-- <el-table-column width="100px" label="Importance">
        <template slot-scope="scope">
          <svg-icon v-for="n in +scope.row.importance" :key="n" icon-class="star" class="meta-item__icon" />
        </template>
      </el-table-column> -->

      <el-table-column class-name="status-col" :label="$t('table.status')" width="110">
        <template slot-scope="scope">
          <el-tag :type="scope.row.published | statusFilter">
            {{ scope.row.published ? $t('article.published') : $t('table.draft') }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" :label="$t('article.view')" width="100">
        <template slot-scope="scope">
          <span>{{ scope.row.view }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" :label="$t('table.actions')" width="240">
        <template slot-scope="scope">
          <router-link :to="'/article/edit/'+scope.row.id">
            <el-button type="primary" size="small">
              {{ $t('table.edit') }}
            </el-button>
          </router-link>

          <el-popconfirm v-if="scope.row.published" placement="top" title="确定下架文章吗？" @onConfirm="articlePublishHandle(scope.row)">
            <el-button slot="reference" type="warning" size="small" style="margin-left: 10px;">{{ $t('article.unpublish') }}</el-button>
          </el-popconfirm>
          <el-popconfirm v-else placement="top" title="确定发布文章吗？" @onConfirm="articlePublishHandle(scope.row)">
            <el-button slot="reference" type="success" size="small" style="margin-left: 10px;">{{ $t('table.publish') }}</el-button>
          </el-popconfirm>

          <el-popconfirm placement="top" title="确定删除该文章吗？" @onConfirm="handleDelete(scope.row)">
            <el-button slot="reference" type="danger" size="small" style="margin-left: 10px;">{{ $t('table.delete') }}</el-button>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.page_size" @pagination="getList" />
  </div>
</template>

<script>
import { fetchList, publishArticle, unpublishArticle, deleteArticle } from '@/api/article'
import { fetchSelectList } from '@/api/tag'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
import waves from '@/directive/waves' // waves directive

export default {
  name: 'ArticleList',
  components: { Pagination },
  directives: { waves },
  filters: {
    statusFilter(status) {
      const statusMap = {
        1: 'success',
        0: 'info'
      }
      return statusMap[status || 0]
    }
  },
  data() {
    return {
      list: null,
      total: 0,
      listLoading: true,
      tagListOptions: [],
      listQuery: {
        page: 1,
        page_size: 10,
        tags: [],
        sort: 'created_at desc'
      }
    }
  },
  created() {
    this.getList()
    this.getTagList()
  },
  methods: {
    // TODO: 确认显示字段，及编辑操作
    getList() {
      this.listLoading = true
      const query = Object.assign({}, this.listQuery)
      if (query.tags.length > 0) {
        query.tags = query.tags.join(',')
      }
      fetchList(query).then(response => {
        this.list = response.data.list
        this.total = response.data.pagination.total
        this.listLoading = false
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
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    },
    getTagNames(tags) {
      tags = tags || []
      const names = []
      tags.forEach(v => {
        names.push(v.name)
      })
      return names.join(',')
    },
    articlePublishHandle(recode) {
      if (recode.published) {
        unpublishArticle(recode.id).then(() => {
          recode.published = 0
          this.$message({
            message: this.$t('common.unpublishSuccess'),
            type: 'success',
            duration: 1000
          })
        })
      } else {
        publishArticle(recode.id).then(() => {
          recode.published = 1
          this.$message({
            message: this.$t('common.publishSuccess'),
            type: 'success',
            duration: 1000
          })
        })
      }
    },
    handleDelete(row) {
      deleteArticle(row.id).then(() => {
        this.getList()
        this.$notify({
          title: this.$t('common.success'),
          message: this.$t('common.deleteSuccess'),
          type: 'success',
          duration: 2000
        })
      })
    }
  }
}
</script>

<style scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
