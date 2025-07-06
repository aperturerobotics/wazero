package fsapi

import experimentalsys "github.com/tetratelabs/wazero/experimental/sys"

// File is a type alias for experimentalsys.PollableFile, which extends
// experimentalsys.File with polling and non-blocking support.
//
// This alias exists so that internal code can continue to reference
// fsapi.File without change, while the canonical definition lives in
// the experimental/sys package for external implementors.
type File = experimentalsys.PollableFile
