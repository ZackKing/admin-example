<?php

namespace App\Models;

use App\Components\Helper;
use Medoo\Medoo;

class Model
{
    private static $_instances = [];
    protected static $_connections = [];
    protected static $_conf = [];

    protected $connection = 'default';
    protected $table = '';
    protected $fields = [];

    protected function __construct()
    {
        self::$_conf = config('database.db');
    }

    /**
     * Singleton Pattern
     * @return static
     */
    public static function instance()
    {
        $class = get_called_class();
        if (!isset(self::$_instances[$class])) {
            self::$_instances[$class] = new static();
        }
        return self::$_instances[$class];
    }

    public function db(string $mode = 'read'): Medoo
    {
        if (empty(self::$_conf[$this->connection]['read'])) { // not write/read mode
            $mode = '';
        }
        $ckey = rtrim($this->connection . '_' . $mode, '_');
        if (empty(self::$_connections[$ckey])) {
            self::$_connections[$ckey] = $this->_connect($this->connection, $mode);
        }
        return self::$_connections[$ckey];
    }

    public function setConnection(string $conn = 'default'): self
    {
        $this->connection = $conn;
        return $this;
    }

    public function getColumns(): array
    {
        return array_keys($this->fields);
    }

    public function insert(array $data): int
    {
        $data = $this->_columnsFilter($data);
        if (!$data) {
            return 0;
        }
        $this->db('write')->insert($this->table, $data);
        return $this->db('write')->id();
    }

    public function multiInsert(array $list, bool $throw = false): int
    {
        if (!$list) {
            return 0;
        }
        foreach ($list as $k => $v) {
            $list[$k] = $this->_columnsFilter($v);
            if (!$list[$k]) { // have empty item
                return 0;
            }
        }
        $pdos = $this->db('write')->insert($this->table, $list);
        if ($this->db('write')->pdo->errorCode() != '00000' && $throw) {
            $error = json_encode($this->db('write')->error);
            throw new \Exception("Db query Error: $error");
        }
        return $pdos ? $pdos->rowCount() : 0;
    }

    public function update(array $where, array $data): int
    {
        $data = $this->_columnsFilter($data);
        if (!$where || !$data) { // safe update
            return 0;
        }
        $pdos = $this->db('write')->update($this->table, $data, $where);
        if (!$pdos || $pdos->errorCode() != '00000') {
            $msg = 'Unknow Db update error!';
            if ($pdos) {
                isset($pdos->errorInfo()[2]) && $msg = $pdos->errorInfo()[2];
            }
            throw new \Exception("Db Update Error: $msg");
        }
        return $pdos ? $pdos->rowCount() : 0;
    }

    public function delete(array $where): int
    {
        if (!$where) { // safe delete
            return 0;
        }
        $pdos = $this->db('write')->delete($this->table, $where);
        return $pdos ? $pdos->rowCount() : 0;
    }

    public function take(array $where = [], array $columns = [], string $mode = 'read'): array
    {
        !$columns && $columns = $this->getColumns();
        if (!$columns) {
            return [];
        }
        $list = $this->db($mode)->select($this->table, $columns, $where);
        return $list ? $this->_settypes($list) : [];
    }

    public function takeFirst(array $where = [], array $columns = [], string $mode = 'read'): array
    {
        !$columns && $columns = $this->getColumns();
        if (!$columns) {
            return [];
        }
        $row = $this->db($mode)->get($this->table, $columns, $where);
        return $row ? $this->_settype($row) : [];
    }

    public function takeWithTotal(array $where, array $columns = []): array
    {
        $limit = [];
        if (!empty($where['LIMIT'])) {
            $limit = $where['LIMIT'];
            unset($where['LIMIT']);
        }
        $data = [
            'list' => [],
            'total' => $this->count($where),
        ];
        if ($data['total'] > 0) {
            !$columns && $columns = $this->getColumns();
            $limit && $where['LIMIT'] = $limit;
            $data['list'] = $this->take($where, $columns);
        }
        return $data;
    }

    public function count(array $where, string $mod = 'read'): int
    {
        return $this->db($mod)->count($this->table, $where);
    }

    public function query(string $sql, array $params = [], bool $throw = false, string $mod = 'read'): array
    {
        $query = $this->db($mod)->query($sql, $params)->fetchAll(\PDO::FETCH_ASSOC);
        if ($this->db($mod)->pdo->errorCode() != '00000' && $throw) {
            $error = json_encode($this->db()->error);
            throw new \Exception("Db query Error: $error");
        }
        return $query ?: [];
    }

    public function queryFirst(string $sql, array $params = [], bool $throw = false, string $mod = 'read'): array
    {
        $query = $this->db($mod)->query($sql, $params)->fetch(\PDO::FETCH_ASSOC);
        if ($this->db($mod)->pdo->errorCode() != '00000' && $throw) {
            $error = json_encode($this->db()->error);
            throw new \Exception("Db query Error: $error");
        }
        return $query ?: [];
    }

    public function lastSql(string $mod = 'read')
    {
        return $this->db($mod)->last();
    }

    public function getTableName()
    {
        return $this->table;
    }

    protected function _settype(array $data): array
    {
        foreach ($data as $k => $v) {
            if (!isset($this->fields[$k])) {
                continue;
            }
            $type = strtolower($this->fields[$k]);
            switch ($type) {
                case 'json':
                    $data[$k] = json_decode($v, true);
                    break;
                default:
                    \settype($data[$k], $type);
                    break;
            }
        }
        return $data;
    }

    protected function _settypes(array $list): array
    {
        foreach ($list as $k => $v) {
            $list[$k] = $this->_settype($v);
        }
        return $list;
    }

    protected function _columnsFilter(array $data): array
    {
        return Helper::fieldFilter($data, $this->getColumns());
    }

    protected function _connect(string $key, string $mode = ''): Medoo
    {
        $conf = [];
        if ($mode) {
            if (empty(self::$_conf[$key][$mode])) {
                throw new \Exception("Db Config Error! Not Found Connection: $key $mode ");
            }
            $conf = array_merge(self::$_conf[$key][$mode], self::$_conf[$key]);
            unset($conf[$mode]);
        } else {
            if (empty(self::$_conf[$key])) {
                throw new \Exception("Db Config Error! Not Found Connection: $key");
            }
            $conf = self::$_conf[$key];
        }
        $connection = new Medoo($conf);
        if ($error = $connection->error) {
            throw new \Exception("Db Conncet $key $mode Error: $error[0] $error[1]");
        }
        return $connection;
    }

}
