function getSignal() {
  const controller = new AbortController()
  let signal = controller.signal
  return signal
}

const config = {
  signal: getSignal(),
}

export default config