import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: 'articles',
    method: 'get',
    params: query
  })
}

export function fetchArticle(id) {
  return request({
    url: 'articles/' + id,
    method: 'get'
  })
}

export function createArticle(data) {
  return request({
    url: 'articles',
    method: 'post',
    data
  })
}

export function updateArticle(id, data) {
  return request({
    url: 'articles/' + id,
    method: 'put',
    data
  })
}

export function publishArticle(id) {
  return request({
    url: 'articles/' + id + '/publish',
    method: 'put'
  })
}

export function unpublishArticle(id) {
  return request({
    url: 'articles/' + id + '/unpublish',
    method: 'put'
  })
}
