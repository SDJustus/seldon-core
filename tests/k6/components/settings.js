function schedulerEndpoint() {
    if (__ENV.SCHEDULER_ENDPOINT) {
        return __ENV.SCHEDULER_ENDPOINT
    }
    return "0.0.0.0:9004"
}

function inferGrpcEndpoint() {
    if (__ENV.INFER_GRPC_ENDPOINT) {
        return __ENV.INFER_GRPC_ENDPOINT
    }
    return "0.0.0.0:9000"
}

function inferHttpEndpoint() {
    if (__ENV.INFER_HTTP_ENDPOINT) {
        return __ENV.INFER_HTTP_ENDPOINT
    }
    return "http://0.0.0.0:9000"
}

function inferHttpIterations() {
    if (__ENV.INFER_HTTP_ITERATIONS) {
        return __ENV.INFER_HTTP_ITERATIONS
    }
    return 1
}

function inferGrpcIterations() {
    if (__ENV.INFER_GRPC_ITERATIONS) {
        return __ENV.INFER_GRPC_ITERATIONS
    }
    return 1
}

function modelType() {
    if (__ENV.MODEL_TYPE) {
        return __ENV.MODEL_TYPE
    }
    return "iris"
}

function loadModel() {
    if (__ENV.SKIP_LOAD_MODEL) {
        return false
    }
    return true
}

function infer() {
    if (__ENV.SKIP_INFER) {
        return false
    }
    return true
}

function unloadModel() {
    if (__ENV.SKIP_UNLOAD_MODEL) {
        return false
    }
    return true
}

function maxNumModels() {
    if (__ENV.MAX_NUM_MODELS) {
        return __ENV.MAX_NUM_MODELS
    }
    return 10
}

function isSchedulerProxy() {
    if (__ENV.SCHEDULER_PROXY) {
        return __ENV.SCHEDULER_PROXY
    }
    return false
}

function isEnvoy() {
    if (__ENV.ENVOY) {
        return (__ENV.ENVOY === "true")
    }
    return true
}

function modelMemoryBytes() {
    if (__ENV.MODEL_MEMORY_BYTES) {
        return __ENV.MODEL_MEMORY_BYTES
    }
    return null
}


export function getConfig() {
    return {
        "schedulerEndpoint": schedulerEndpoint(),
        "inferHttpEndpoint": inferHttpEndpoint(),
        "inferGrpcEndpoint": inferGrpcEndpoint(),
        "inferHttpIterations": inferHttpIterations(),
        "inferGrpcIterations": inferGrpcIterations(),
        "modelType": modelType(),
        "loadModel": loadModel(),
        "infer": infer(),
        "unloadModel": unloadModel(),
        "maxNumModels": maxNumModels(),
        "isSchedulerProxy": isSchedulerProxy(),
        "isEnvoy": isEnvoy(),
        "modelMemoryBytes": modelMemoryBytes()
    }
}