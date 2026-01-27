/// <reference types="vite/client" />
/// <reference types="vite-svg-loader" />

import type { MessageApiInjection,DialogApiInjection,NotificationApiInjection } from "naive-ui/es/message/src/MessageProvider";



declare global{
  interface Window {
    $message: MessageApiInjection
    $dialog: DialogApiInjection
    $notification: NotificationApiInjection
  }
}
