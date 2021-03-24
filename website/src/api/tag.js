import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/tags',
    method: 'get',
    params: query
  })
}

export function fetchSelectList(query) {
  return request({
    url: '/simple_tags',
    method: 'get',
    params: query
  })
}

export function fetchTag(id) {
  return request({
    url: '/tags/' + id,
    method: 'get'
  })
}

export function createTag(data) {
  return request({
    url: '/tags',
    method: 'post',
    data
  })
}

export function updateTag(id, data) {
  return request({
    url: '/tags/' + id,
    method: 'put',
    data
  })
}

export function deleteTag(id) {
  return request({
    url: '/tags/' + id,
    method: 'delete'
  })
}
