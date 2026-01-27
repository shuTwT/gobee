import { config, XSSPlugin } from 'md-editor-v3'
import type markdownit from 'markdown-it'

const ProductCardPlugin = (md: markdownit, options: {}) => {
  md.block.ruler.before('fence','product-card', (state, startLine, endLine, slient) => {
    // 检查当前行是否以```product-card开头
    const src = state.src
    const pos = state.bMarks[startLine] + state.tShift[startLine]
    const max = state.eMarks[startLine]

    // 匹配代码块起始
    if (src.slice(pos, pos + 15) !== '```product-card') {
      return false
    }

    // 查找代码块结束
    let endPos = -1
    for (let i = startLine + 1; i < endLine; i++) {
      const linePos = state.bMarks[i] + state.tShift[i]
      const lineMax = state.eMarks[i]

      if (src.slice(linePos, linePos + 3) === '```') {
        endPos = i
        break
      }
    }

    if (endPos === -1) {
      return false
    }

    // 提取代码块内容
    const contentStart = startLine + 1
    const contentEnd = endPos
    const content = src.slice(
      state.bMarks[contentStart] + state.tShift[contentStart],
      state.bMarks[contentEnd] + state.tShift[contentEnd],
    )

    // 解析键值对 - 直接构建二维数组
    const attributes: [string, string][] = []
    content.split('\n').forEach((line) => {
      const [key, ...valueParts] = line.trim().split(' ')
      if (key && valueParts.length > 0) {
        attributes.push([key, valueParts.join(' ')])
      }
    })

    // 创建token
    state.line = endPos + 1
    const token = state.push('product-card', 'product-card', 0)
    token.attrs = attributes // 直接赋值二维数组

    return true
  })

  md.renderer.rules['product-card'] = (tokens, idx, ops, env, slf) => {
    const token = tokens[idx]
    const attrs = token.attrs || []
    console.log(Object.fromEntries(attrs))
    const attrObj = Object.fromEntries(attrs)
    let attrString = ``
    for(const [key, value] of attrs){
      attrString += `${key}="${value}" `
    }
    const code = token.content
    console.log(code)

    return `<product-card ${attrString}></product-card>`

  }
}

const editorConfig = config({
  markdownItPlugins(plugins) {
    return [
      ...plugins,
      {
        type: 'product-card',
        plugin: ProductCardPlugin,
        options: {},
      },
    ]
  },
})

export { editorConfig }
