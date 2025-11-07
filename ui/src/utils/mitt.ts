import type { Emitter } from "mitt";
import mitt from "mitt"

type Events = {
  logoChange:boolean
}

export const emitter:Emitter<Events> = mitt<Events>()
