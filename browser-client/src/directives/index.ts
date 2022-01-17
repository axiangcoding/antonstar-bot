import injectThemes from './injectThemes.ts'

const directives: any = {
	injectThemes,
}

export default function injectDirectives(app: any) {
  for (let i in directives) {
		directives[i](app)
	}
}
