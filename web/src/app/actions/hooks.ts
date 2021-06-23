import { urlParamSelector, urlPathSelector } from '../selectors'
import { setURLParam as setParam, resetURLParams as resetParams } from './main'
import { useSelector, useDispatch } from 'react-redux'
import { useLocation } from 'react-router'
import { warn } from '../util/debug'

export type Value = string | boolean | number | string[]

export function useURLParam<T extends Value>(
  name: string,
  defaultValue: T,
): [T, (newValue: T) => void] {
  const dispatch = useDispatch()
  const urlParam = useSelector(urlParamSelector)
  const urlPath = useSelector(urlPathSelector)
  const value = urlParam(name, defaultValue) as T

  function setValue(newValue: T): void {
    if (window.location.pathname !== urlPath) {
      warn(
        'useURLParam was called to set a parameter, but location.pathname has changed, aborting',
      )
      return
    }
    dispatch(setParam(name, newValue, defaultValue))
  }

  return [value, setValue]
}

export function useResetURLParams(...keys: Array<string>): () => void {
  const dispatch = useDispatch()
  const urlPath = useSelector(urlPathSelector)
  function resetURLParams(): void {
    if (window.location.pathname !== urlPath) {
      warn(
        'useResetURLParams was called to reset parameters, but location.pathname has changed, aborting',
      )
      return
    }
    dispatch(resetParams(...keys))
  }

  return resetURLParams
}

// useURLKey provides a unique identifier per page
// e.g. location.key as provided by React Router
export function useURLKey(): string {
  const l = useLocation()
  return l.key ?? ''
}
