import request from '@/utils/request'

export function getNewVisitCount() {
  return request({
    url: '/statistics/new_visit_count',
    method: 'get'
  })
}

export function getVisitCount() {
  return request({
    url: '/statistics/visit_count',
    method: 'get'
  })
}

export function getViewCount() {
  return request({
    url: '/statistics/view_count',
    method: 'get'
  })
}

export function getArticlyCount() {
  return request({
    url: '/statistics/articly_count',
    method: 'get'
  })
}

export function getStatisticsRange() {
  return request({
    url: '/statistics/range',
    method: 'get'
  })
}
