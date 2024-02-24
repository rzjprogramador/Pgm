export interface Debugger {
  via_console: DebuggarViaConsole

}

interface DebuggarViaConsole {
  observacoes: string
  em_pasta_personal: CompilarConsole
  no_mesmo_local: CompilarConsole
}

interface CompilarConsole {
  compilar: string
  mostrar_no_console: string
  detalhes: string

}