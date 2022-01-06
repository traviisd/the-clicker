interface go {
  "main": {
    "App": {
		Clicks():Promise<string>
		ClicksPerSecond():Promise<string>
		SetDelay(arg1:string):Promise<void>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
