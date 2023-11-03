<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/NodeType.proto

namespace NodeType;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>NodeType.Complexity</code>
 */
class Complexity extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>optional int32 cyclomatic = 1;</code>
     */
    protected $cyclomatic = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $cyclomatic
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\NodeType::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>optional int32 cyclomatic = 1;</code>
     * @return int
     */
    public function getCyclomatic()
    {
        return isset($this->cyclomatic) ? $this->cyclomatic : 0;
    }

    public function hasCyclomatic()
    {
        return isset($this->cyclomatic);
    }

    public function clearCyclomatic()
    {
        unset($this->cyclomatic);
    }

    /**
     * Generated from protobuf field <code>optional int32 cyclomatic = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setCyclomatic($var)
    {
        GPBUtil::checkInt32($var);
        $this->cyclomatic = $var;

        return $this;
    }

}
