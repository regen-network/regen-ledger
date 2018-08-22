package ceres.rocksdb

import ceres.storage.IKVStore
import org.rocksdb.Options
import org.rocksdb.RocksDB

fun openRocksDb(path: String): RocksDB {
    val options = Options().setCreateIfMissing(true)
    val db = RocksDB.open(options, path)
    if(db == null)
        throw IllegalStateException()
    return db
}

class RocksDBKVStore(val db: RocksDB): IKVStore<ByteArray, ByteArray> {
    companion object {
        init {
            RocksDB.loadLibrary()
        }
    }

    constructor(path: String): this(openRocksDb(path)) {}

    override suspend fun get(key: ByteArray): ByteArray? = db.get(key)

    override suspend fun set(key: ByteArray, value: ByteArray) = db.put(key, value)

    override suspend fun delete(key: ByteArray) = db.delete(key)
}
