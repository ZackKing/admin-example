
export default {
  mounted(el, binding) {
    const text = binding.value || 'Watermark'
    const fontSize = 25
    const padding = 50

    const tempCanvas = document.createElement('canvas')
    const tempCtx = tempCanvas.getContext('2d')
    tempCtx.font = `${fontSize}px Arial`
    const textWidth = tempCtx.measureText(text).width

    const canvasWidth = textWidth + padding * 2
    const canvasHeight = canvasWidth

    const canvas = document.createElement('canvas')
    canvas.width = canvasWidth
    canvas.height = canvasHeight

    const ctx = canvas.getContext('2d')
    ctx.clearRect(0, 0, canvas.width, canvas.height)
    ctx.font = `${fontSize}px Arial`
    ctx.fillStyle = 'rgba(200, 200, 200, 0.20)'
    ctx.textAlign = 'center'
    ctx.textBaseline = 'middle'
    ctx.translate(canvas.width / 2, canvas.height / 2)
    ctx.rotate(-Math.PI / 4)
    ctx.fillText(text, 0, 0)

    const watermarkDiv = document.createElement('div')
    watermarkDiv.style.position = 'absolute'
    watermarkDiv.style.top = '0'
    watermarkDiv.style.left = '0'
    watermarkDiv.style.width = '100%'
    watermarkDiv.style.height = '100%'
    watermarkDiv.style.pointerEvents = 'none'
    watermarkDiv.style.zIndex = '9999'
    watermarkDiv.style.backgroundImage = `url(${canvas.toDataURL('image/png')})`

    el.style.position = 'relative'
    el.appendChild(watermarkDiv)
  }
}
