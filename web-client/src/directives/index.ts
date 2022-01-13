import { adhereToTheTop } from './adhereToTheTop.ts'

const directives = {
  adhereToTheTop
}

export default function injectDirectives(app: any) {
  for (let i in directives) {
    directives[i](app)
  }
}
