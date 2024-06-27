(module
  (type (;0;) (func (param i32 i32 i32 i32) (result i32)))
  (type (;1;) (func (param i32 i32) (result i32)))
  (type (;2;) (func (param i32)))
  (type (;3;) (func (param i32 i32)))
  (type (;4;) (func))
  (type (;5;) (func (param i32 i32 i32) (result i32)))
  (type (;6;) (func (param i32) (result i32)))
  (type (;7;) (func (result i32)))
  (type (;8;) (func (param i32 i32 i32 i32 i32 i32) (result i32)))
  (type (;9;) (func (param i32 i32 i32 i32 i32)))
  (import "wasi_snapshot_preview1" "args_get" (func (;0;) (type 1)))
  (import "wasi_snapshot_preview1" "args_sizes_get" (func (;1;) (type 1)))
  (import "wasi_snapshot_preview1" "proc_exit" (func (;2;) (type 2)))
  (import "env" "set" (func (;3;) (type 3)))
  (func (;4;) (type 4))
  (func (;5;) (type 4)
    (local i32)
    block  ;; label = @1
      block  ;; label = @2
        i32.const 0
        i32.load offset=16778288
        br_if 0 (;@2;)
        i32.const 0
        i32.const 1
        i32.store offset=16778288
        call 4
        call 13
        local.set 0
        call 20
        local.get 0
        br_if 1 (;@1;)
        return
      end
      unreachable
      unreachable
    end
    local.get 0
    call 16
    unreachable)
  (func (;6;) (type 1) (param i32 i32) (result i32)
    (local i32 i32 i32 i64 i64 i32 i32 i32)
    i32.const 0
    local.set 2
    block  ;; label = @1
      local.get 0
      local.get 0
      i32.const -1
      i32.add
      i32.and
      br_if 0 (;@1;)
      local.get 1
      i32.const -57
      i32.gt_u
      br_if 0 (;@1;)
      loop  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 1
            i32.const 3
            i32.add
            i32.const -4
            i32.and
            i32.const 8
            local.get 1
            i32.const 8
            i32.gt_u
            select
            local.tee 1
            i32.const 127
            i32.gt_u
            br_if 0 (;@4;)
            local.get 1
            i32.const 3
            i32.shr_u
            i32.const -1
            i32.add
            local.set 3
            br 1 (;@3;)
          end
          local.get 1
          i32.clz
          local.set 4
          block  ;; label = @4
            local.get 1
            i32.const 4095
            i32.gt_u
            br_if 0 (;@4;)
            local.get 1
            i32.const 29
            local.get 4
            i32.sub
            i32.shr_u
            i32.const 4
            i32.xor
            local.get 4
            i32.const 2
            i32.shl
            i32.sub
            i32.const 110
            i32.add
            local.set 3
            br 1 (;@3;)
          end
          local.get 1
          i32.const 30
          local.get 4
          i32.sub
          i32.shr_u
          i32.const 2
          i32.xor
          local.get 4
          i32.const 1
          i32.shl
          i32.sub
          i32.const 71
          i32.add
          local.tee 4
          i32.const 63
          local.get 4
          i32.const 63
          i32.lt_u
          select
          local.set 3
        end
        local.get 0
        i32.const 16
        local.get 0
        i32.const 16
        i32.gt_u
        select
        local.set 0
        block  ;; label = @3
          i32.const 0
          i64.load offset=16778296
          local.tee 5
          local.get 3
          i64.extend_i32_u
          i64.shr_u
          local.tee 6
          i64.eqz
          br_if 0 (;@3;)
          loop  ;; label = @4
            local.get 6
            local.get 6
            i64.ctz
            local.tee 5
            i64.shr_u
            local.set 6
            block  ;; label = @5
              block  ;; label = @6
                local.get 3
                local.get 5
                i32.wrap_i64
                i32.add
                local.tee 3
                i32.const 4
                i32.shl
                local.tee 7
                i32.const 16777272
                i32.add
                i32.load
                local.tee 4
                local.get 7
                i32.const 16777264
                i32.add
                local.tee 8
                i32.eq
                br_if 0 (;@6;)
                local.get 4
                local.get 0
                local.get 1
                call 7
                local.tee 2
                br_if 5 (;@1;)
                local.get 4
                i32.load offset=4
                local.get 4
                i32.load offset=8
                local.tee 2
                i32.store offset=8
                local.get 2
                local.get 4
                i32.load offset=4
                i32.store offset=4
                local.get 4
                local.get 8
                i32.store offset=8
                local.get 4
                local.get 7
                i32.const 16777268
                i32.add
                local.tee 7
                i32.load
                i32.store offset=4
                local.get 7
                local.get 4
                i32.store
                local.get 4
                i32.load offset=4
                local.get 4
                i32.store offset=8
                local.get 6
                i64.const 1
                i64.shr_u
                local.set 6
                local.get 3
                i32.const 1
                i32.add
                local.set 3
                br 1 (;@5;)
              end
              i32.const 0
              i32.const 0
              i64.load offset=16778296
              i64.const -2
              local.get 3
              i64.extend_i32_u
              i64.rotl
              i64.and
              i64.store offset=16778296
              local.get 6
              i64.const 1
              i64.xor
              local.set 6
            end
            local.get 6
            i64.const 0
            i64.ne
            br_if 0 (;@4;)
          end
          i32.const 0
          i64.load offset=16778296
          local.set 5
        end
        i32.const 63
        local.get 5
        i64.clz
        i32.wrap_i64
        i32.sub
        local.set 8
        block  ;; label = @3
          block  ;; label = @4
            local.get 5
            i64.eqz
            i32.eqz
            br_if 0 (;@4;)
            i32.const 0
            local.set 4
            br 1 (;@3;)
          end
          local.get 8
          i32.const 4
          i32.shl
          local.tee 3
          i32.const 16777272
          i32.add
          i32.load
          local.set 4
          local.get 5
          i64.const 1073741824
          i64.lt_u
          br_if 0 (;@3;)
          local.get 4
          local.get 3
          i32.const 16777264
          i32.add
          local.tee 7
          i32.eq
          br_if 0 (;@3;)
          i32.const -100
          local.set 3
          loop  ;; label = @4
            local.get 3
            i32.const 1
            i32.add
            local.tee 3
            i32.eqz
            br_if 1 (;@3;)
            local.get 4
            local.get 0
            local.get 1
            call 7
            local.tee 2
            br_if 3 (;@1;)
            local.get 4
            i32.load offset=8
            local.tee 4
            local.get 7
            i32.ne
            br_if 0 (;@4;)
          end
          local.get 7
          local.set 4
        end
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              i32.const 0
              i32.load offset=16778304
              br_if 0 (;@5;)
              local.get 1
              i32.const 48
              i32.add
              local.set 2
              i32.const 16842752
              local.set 7
              block  ;; label = @6
                i32.const 16842752
                br_if 0 (;@6;)
                i32.const 0
                call 18
                local.set 7
              end
              i32.const 16778496
              local.set 3
              local.get 7
              i32.const 16778496
              i32.sub
              local.get 2
              i32.ge_u
              br_if 1 (;@4;)
            end
            local.get 1
            i32.const 65583
            i32.add
            i32.const -65536
            i32.and
            local.tee 7
            call 18
            local.tee 3
            i32.const -1
            i32.eq
            br_if 1 (;@3;)
            local.get 3
            local.get 7
            i32.add
            local.set 7
          end
          local.get 7
          i32.const -4
          i32.add
          i32.const 16
          i32.store
          local.get 7
          i32.const -16
          i32.add
          local.tee 2
          i32.const 16
          i32.store
          i32.const 0
          local.set 4
          block  ;; label = @4
            i32.const 0
            i32.load offset=16778304
            local.tee 8
            i32.eqz
            br_if 0 (;@4;)
            local.get 8
            i32.load offset=8
            local.set 4
          end
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 3
                local.get 4
                i32.ne
                br_if 0 (;@6;)
                local.get 3
                local.get 3
                i32.const -4
                i32.add
                i32.load
                i32.const -2
                i32.and
                i32.sub
                local.tee 4
                i32.const -4
                i32.add
                i32.load
                local.set 9
                local.get 8
                local.get 7
                i32.store offset=8
                block  ;; label = @7
                  local.get 4
                  local.get 9
                  i32.const -2
                  i32.and
                  i32.sub
                  local.tee 4
                  local.get 4
                  i32.load
                  i32.add
                  i32.const -4
                  i32.add
                  i32.load8_u
                  i32.const 1
                  i32.and
                  i32.eqz
                  br_if 0 (;@7;)
                  local.get 4
                  i32.load offset=4
                  local.get 4
                  i32.load offset=8
                  local.tee 3
                  i32.store offset=8
                  local.get 3
                  local.get 4
                  i32.load offset=4
                  i32.store offset=4
                  local.get 4
                  local.get 2
                  local.get 4
                  i32.sub
                  local.tee 3
                  i32.store
                  local.get 4
                  local.get 3
                  i32.const -4
                  i32.and
                  i32.add
                  i32.const -4
                  i32.add
                  local.get 3
                  i32.const 1
                  i32.or
                  i32.store
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 4
                      i32.load
                      i32.const -8
                      i32.add
                      local.tee 3
                      i32.const 127
                      i32.gt_u
                      br_if 0 (;@9;)
                      local.get 3
                      i32.const 3
                      i32.shr_u
                      i32.const -1
                      i32.add
                      local.set 3
                      br 1 (;@8;)
                    end
                    local.get 3
                    i32.clz
                    local.set 7
                    block  ;; label = @9
                      local.get 3
                      i32.const 4095
                      i32.gt_u
                      br_if 0 (;@9;)
                      local.get 3
                      i32.const 29
                      local.get 7
                      i32.sub
                      i32.shr_u
                      i32.const 4
                      i32.xor
                      local.get 7
                      i32.const 2
                      i32.shl
                      i32.sub
                      i32.const 110
                      i32.add
                      local.set 3
                      br 1 (;@8;)
                    end
                    local.get 3
                    i32.const 30
                    local.get 7
                    i32.sub
                    i32.shr_u
                    i32.const 2
                    i32.xor
                    local.get 7
                    i32.const 1
                    i32.shl
                    i32.sub
                    i32.const 71
                    i32.add
                    local.tee 3
                    i32.const 63
                    local.get 3
                    i32.const 63
                    i32.lt_u
                    select
                    local.set 3
                  end
                  local.get 4
                  local.get 3
                  i32.const 4
                  i32.shl
                  local.tee 7
                  i32.const 16777264
                  i32.add
                  i32.store offset=4
                  local.get 4
                  local.get 7
                  i32.const 16777272
                  i32.add
                  local.tee 7
                  i32.load
                  i32.store offset=8
                  local.get 7
                  local.get 4
                  i32.store
                  br 3 (;@4;)
                end
                local.get 3
                i32.const -16
                i32.add
                local.set 4
                br 1 (;@5;)
              end
              local.get 3
              i32.const 16
              i32.store
              local.get 3
              i32.const 12
              i32.add
              i32.const 16
              i32.store
              local.get 3
              local.get 7
              i32.store offset=8
              local.get 3
              local.get 8
              i32.store offset=4
              i32.const 0
              local.get 3
              i32.store offset=16778304
              local.get 3
              i32.const 16
              i32.add
              local.set 4
            end
            local.get 4
            local.get 2
            local.get 4
            i32.sub
            local.tee 3
            i32.store
            local.get 4
            local.get 3
            i32.const -4
            i32.and
            i32.add
            i32.const -4
            i32.add
            local.get 3
            i32.const 1
            i32.or
            i32.store
            block  ;; label = @5
              block  ;; label = @6
                local.get 4
                i32.load
                i32.const -8
                i32.add
                local.tee 3
                i32.const 127
                i32.gt_u
                br_if 0 (;@6;)
                local.get 3
                i32.const 3
                i32.shr_u
                i32.const -1
                i32.add
                local.set 3
                br 1 (;@5;)
              end
              local.get 3
              i32.clz
              local.set 7
              block  ;; label = @6
                local.get 3
                i32.const 4095
                i32.gt_u
                br_if 0 (;@6;)
                local.get 3
                i32.const 29
                local.get 7
                i32.sub
                i32.shr_u
                i32.const 4
                i32.xor
                local.get 7
                i32.const 2
                i32.shl
                i32.sub
                i32.const 110
                i32.add
                local.set 3
                br 1 (;@5;)
              end
              local.get 3
              i32.const 30
              local.get 7
              i32.sub
              i32.shr_u
              i32.const 2
              i32.xor
              local.get 7
              i32.const 1
              i32.shl
              i32.sub
              i32.const 71
              i32.add
              local.tee 3
              i32.const 63
              local.get 3
              i32.const 63
              i32.lt_u
              select
              local.set 3
            end
            local.get 4
            local.get 3
            i32.const 4
            i32.shl
            local.tee 7
            i32.const 16777264
            i32.add
            i32.store offset=4
            local.get 4
            local.get 7
            i32.const 16777272
            i32.add
            local.tee 7
            i32.load
            i32.store offset=8
            local.get 7
            local.get 4
            i32.store
          end
          local.get 4
          i32.load offset=8
          local.get 4
          i32.store offset=4
          i32.const 0
          local.set 2
          i32.const 0
          i32.const 0
          i64.load offset=16778296
          i64.const 1
          local.get 3
          i64.extend_i32_u
          i64.shl
          i64.or
          i64.store offset=16778296
          local.get 0
          local.get 0
          i32.const -1
          i32.add
          i32.and
          br_if 2 (;@1;)
          local.get 1
          i32.const -57
          i32.le_u
          br_if 1 (;@2;)
          br 2 (;@1;)
        end
      end
      block  ;; label = @2
        local.get 4
        i32.eqz
        br_if 0 (;@2;)
        local.get 4
        local.get 8
        i32.const 4
        i32.shl
        i32.const 16777264
        i32.add
        local.tee 3
        i32.eq
        br_if 0 (;@2;)
        loop  ;; label = @3
          local.get 4
          local.get 0
          local.get 1
          call 7
          local.tee 2
          br_if 2 (;@1;)
          local.get 4
          i32.load offset=8
          local.tee 4
          local.get 3
          i32.ne
          br_if 0 (;@3;)
        end
      end
      i32.const 0
      local.set 2
    end
    local.get 2)
  (func (;7;) (type 5) (param i32 i32 i32) (result i32)
    (local i32 i32 i32)
    i32.const 0
    local.set 3
    block  ;; label = @1
      local.get 1
      local.get 0
      i32.const 4
      i32.add
      local.tee 4
      i32.add
      i32.const -1
      i32.add
      i32.const 0
      local.get 1
      i32.sub
      i32.and
      local.tee 1
      local.get 2
      i32.add
      local.get 0
      local.get 0
      i32.load
      i32.add
      i32.const -4
      i32.add
      i32.gt_u
      br_if 0 (;@1;)
      local.get 0
      i32.load offset=4
      local.get 0
      i32.load offset=8
      local.tee 3
      i32.store offset=8
      local.get 3
      local.get 0
      i32.load offset=4
      i32.store offset=4
      block  ;; label = @2
        block  ;; label = @3
          local.get 4
          local.get 1
          i32.ne
          br_if 0 (;@3;)
          local.get 0
          i32.load
          local.set 1
          br 1 (;@2;)
        end
        local.get 0
        i32.load
        local.set 5
        local.get 0
        local.get 0
        i32.const -4
        i32.add
        i32.load
        i32.const -2
        i32.and
        i32.sub
        local.tee 3
        local.get 3
        i32.load
        local.get 1
        local.get 4
        i32.sub
        local.tee 1
        i32.add
        local.tee 4
        i32.store
        local.get 3
        local.get 4
        i32.const -4
        i32.and
        i32.add
        i32.const -4
        i32.add
        local.get 4
        i32.store
        local.get 0
        local.get 1
        i32.add
        local.tee 0
        local.get 5
        local.get 1
        i32.sub
        local.tee 1
        i32.store
      end
      block  ;; label = @2
        block  ;; label = @3
          local.get 2
          i32.const 24
          i32.add
          local.get 1
          i32.gt_u
          br_if 0 (;@3;)
          local.get 0
          local.get 2
          i32.add
          i32.const 8
          i32.add
          local.tee 3
          local.get 1
          local.get 2
          i32.sub
          i32.const -8
          i32.add
          local.tee 1
          i32.store
          local.get 3
          local.get 1
          i32.const -4
          i32.and
          i32.add
          i32.const -4
          i32.add
          local.get 1
          i32.const 1
          i32.or
          i32.store
          block  ;; label = @4
            block  ;; label = @5
              local.get 3
              i32.load
              i32.const -8
              i32.add
              local.tee 1
              i32.const 127
              i32.gt_u
              br_if 0 (;@5;)
              local.get 1
              i32.const 3
              i32.shr_u
              i32.const -1
              i32.add
              local.set 1
              br 1 (;@4;)
            end
            local.get 1
            i32.clz
            local.set 4
            block  ;; label = @5
              local.get 1
              i32.const 4095
              i32.gt_u
              br_if 0 (;@5;)
              local.get 1
              i32.const 29
              local.get 4
              i32.sub
              i32.shr_u
              i32.const 4
              i32.xor
              local.get 4
              i32.const 2
              i32.shl
              i32.sub
              i32.const 110
              i32.add
              local.set 1
              br 1 (;@4;)
            end
            local.get 1
            i32.const 30
            local.get 4
            i32.sub
            i32.shr_u
            i32.const 2
            i32.xor
            local.get 4
            i32.const 1
            i32.shl
            i32.sub
            i32.const 71
            i32.add
            local.tee 1
            i32.const 63
            local.get 1
            i32.const 63
            i32.lt_u
            select
            local.set 1
          end
          local.get 3
          local.get 1
          i32.const 4
          i32.shl
          local.tee 4
          i32.const 16777264
          i32.add
          i32.store offset=4
          local.get 3
          local.get 4
          i32.const 16777272
          i32.add
          local.tee 4
          i32.load
          i32.store offset=8
          local.get 4
          local.get 3
          i32.store
          local.get 3
          i32.load offset=8
          local.get 3
          i32.store offset=4
          i32.const 0
          i32.const 0
          i64.load offset=16778296
          i64.const 1
          local.get 1
          i64.extend_i32_u
          i64.shl
          i64.or
          i64.store offset=16778296
          local.get 0
          local.get 2
          i32.const 8
          i32.add
          local.tee 1
          i32.store
          local.get 0
          local.get 1
          i32.const -4
          i32.and
          i32.add
          local.set 2
          br 1 (;@2;)
        end
        local.get 0
        local.get 1
        i32.add
        local.set 2
      end
      local.get 2
      i32.const -4
      i32.add
      local.get 1
      i32.store
      local.get 0
      i32.const 4
      i32.add
      local.set 3
    end
    local.get 3)
  (func (;8;) (type 6) (param i32) (result i32)
    i32.const 16
    local.get 0
    call 6)
  (func (;9;) (type 2) (param i32)
    (local i32 i32 i32 i32)
    block  ;; label = @1
      local.get 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      i32.const -4
      i32.add
      local.tee 1
      i32.load
      local.set 2
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          i32.const -8
          i32.add
          i32.load
          local.tee 0
          local.get 0
          i32.const -2
          i32.and
          local.tee 3
          i32.ne
          br_if 0 (;@3;)
          local.get 2
          local.set 3
          local.get 1
          local.set 0
          br 1 (;@2;)
        end
        local.get 1
        local.get 3
        i32.sub
        local.tee 0
        i32.load offset=4
        local.get 0
        i32.load offset=8
        local.tee 4
        i32.store offset=8
        local.get 4
        local.get 0
        i32.load offset=4
        i32.store offset=4
        local.get 3
        local.get 2
        i32.add
        local.set 3
        local.get 1
        i32.load
        local.set 2
      end
      block  ;; label = @2
        local.get 1
        local.get 2
        i32.add
        local.tee 1
        i32.load
        local.tee 2
        local.get 1
        local.get 2
        i32.add
        i32.const -4
        i32.add
        i32.load
        i32.eq
        br_if 0 (;@2;)
        local.get 1
        i32.load offset=4
        local.get 1
        i32.load offset=8
        local.tee 2
        i32.store offset=8
        local.get 2
        local.get 1
        i32.load offset=4
        i32.store offset=4
        local.get 1
        i32.load
        local.get 3
        i32.add
        local.set 3
      end
      local.get 0
      local.get 3
      i32.store
      local.get 0
      local.get 3
      i32.const -4
      i32.and
      i32.add
      i32.const -4
      i32.add
      local.get 3
      i32.const 1
      i32.or
      i32.store
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          i32.load
          i32.const -8
          i32.add
          local.tee 3
          i32.const 127
          i32.gt_u
          br_if 0 (;@3;)
          local.get 3
          i32.const 3
          i32.shr_u
          i32.const -1
          i32.add
          local.set 3
          br 1 (;@2;)
        end
        local.get 3
        i32.clz
        local.set 1
        block  ;; label = @3
          local.get 3
          i32.const 4095
          i32.gt_u
          br_if 0 (;@3;)
          local.get 3
          i32.const 29
          local.get 1
          i32.sub
          i32.shr_u
          i32.const 4
          i32.xor
          local.get 1
          i32.const 2
          i32.shl
          i32.sub
          i32.const 110
          i32.add
          local.set 3
          br 1 (;@2;)
        end
        local.get 3
        i32.const 30
        local.get 1
        i32.sub
        i32.shr_u
        i32.const 2
        i32.xor
        local.get 1
        i32.const 1
        i32.shl
        i32.sub
        i32.const 71
        i32.add
        local.tee 3
        i32.const 63
        local.get 3
        i32.const 63
        i32.lt_u
        select
        local.set 3
      end
      local.get 0
      local.get 3
      i32.const 4
      i32.shl
      local.tee 1
      i32.const 16777264
      i32.add
      i32.store offset=4
      local.get 0
      local.get 1
      i32.const 16777272
      i32.add
      local.tee 1
      i32.load
      i32.store offset=8
      local.get 1
      local.get 0
      i32.store
      local.get 0
      i32.load offset=8
      local.get 0
      i32.store offset=4
      i32.const 0
      i32.const 0
      i64.load offset=16778296
      i64.const 1
      local.get 3
      i64.extend_i32_u
      i64.shl
      i64.or
      i64.store offset=16778296
    end)
  (func (;10;) (type 1) (param i32 i32) (result i32)
    block  ;; label = @1
      i32.const 16
      local.get 1
      local.get 0
      i32.mul
      local.tee 1
      call 6
      local.tee 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      i32.const 0
      local.get 1
      call 21
      drop
    end
    local.get 0)
  (func (;11;) (type 2) (param i32)
    local.get 0
    call 16
    unreachable)
  (func (;12;) (type 1) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call 22)
  (func (;13;) (type 7) (result i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 0
              i32.const 8
              i32.add
              local.get 0
              i32.const 12
              i32.add
              call 15
              br_if 0 (;@5;)
              local.get 0
              i32.load offset=8
              i32.const 1
              i32.add
              local.tee 1
              i32.eqz
              br_if 1 (;@4;)
              local.get 0
              i32.load offset=12
              call 8
              local.tee 2
              i32.eqz
              br_if 2 (;@3;)
              local.get 1
              i32.const 4
              call 10
              local.tee 1
              i32.eqz
              br_if 3 (;@2;)
              local.get 1
              local.get 2
              call 14
              br_if 4 (;@1;)
              local.get 0
              i32.load offset=8
              local.get 1
              call 12
              local.set 1
              local.get 0
              i32.const 16
              i32.add
              global.set 0
              local.get 1
              return
            end
            i32.const 71
            call 11
            unreachable
          end
          i32.const 70
          call 11
          unreachable
        end
        i32.const 70
        call 11
        unreachable
      end
      local.get 2
      call 9
      i32.const 70
      call 11
      unreachable
    end
    local.get 2
    call 9
    local.get 1
    call 9
    i32.const 71
    call 11
    unreachable)
  (func (;14;) (type 1) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call 0
    i32.const 65535
    i32.and)
  (func (;15;) (type 1) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call 1
    i32.const 65535
    i32.and)
  (func (;16;) (type 2) (param i32)
    local.get 0
    call 2
    unreachable)
  (func (;17;) (type 4)
    unreachable
    unreachable)
  (func (;18;) (type 6) (param i32) (result i32)
    block  ;; label = @1
      local.get 0
      br_if 0 (;@1;)
      memory.size
      i32.const 16
      i32.shl
      return
    end
    block  ;; label = @1
      local.get 0
      i32.const 65535
      i32.and
      br_if 0 (;@1;)
      local.get 0
      i32.const -1
      i32.le_s
      br_if 0 (;@1;)
      block  ;; label = @2
        local.get 0
        i32.const 16
        i32.shr_u
        memory.grow
        local.tee 0
        i32.const -1
        i32.ne
        br_if 0 (;@2;)
        i32.const 0
        i32.const 48
        i32.store offset=16778308
        i32.const -1
        return
      end
      local.get 0
      i32.const 16
      i32.shl
      return
    end
    call 17
    unreachable)
  (func (;19;) (type 4))
  (func (;20;) (type 4)
    call 19
    call 19)
  (func (;21;) (type 5) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i64)
    block  ;; label = @1
      local.get 2
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      i32.store8
      local.get 0
      local.get 2
      i32.add
      local.tee 3
      i32.const -1
      i32.add
      local.get 1
      i32.store8
      local.get 2
      i32.const 3
      i32.lt_u
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      i32.store8 offset=2
      local.get 0
      local.get 1
      i32.store8 offset=1
      local.get 3
      i32.const -3
      i32.add
      local.get 1
      i32.store8
      local.get 3
      i32.const -2
      i32.add
      local.get 1
      i32.store8
      local.get 2
      i32.const 7
      i32.lt_u
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      i32.store8 offset=3
      local.get 3
      i32.const -4
      i32.add
      local.get 1
      i32.store8
      local.get 2
      i32.const 9
      i32.lt_u
      br_if 0 (;@1;)
      local.get 0
      i32.const 0
      local.get 0
      i32.sub
      i32.const 3
      i32.and
      local.tee 4
      i32.add
      local.tee 3
      local.get 1
      i32.const 255
      i32.and
      i32.const 16843009
      i32.mul
      local.tee 1
      i32.store
      local.get 3
      local.get 2
      local.get 4
      i32.sub
      i32.const -4
      i32.and
      local.tee 4
      i32.add
      local.tee 2
      i32.const -4
      i32.add
      local.get 1
      i32.store
      local.get 4
      i32.const 9
      i32.lt_u
      br_if 0 (;@1;)
      local.get 3
      local.get 1
      i32.store offset=8
      local.get 3
      local.get 1
      i32.store offset=4
      local.get 2
      i32.const -8
      i32.add
      local.get 1
      i32.store
      local.get 2
      i32.const -12
      i32.add
      local.get 1
      i32.store
      local.get 4
      i32.const 25
      i32.lt_u
      br_if 0 (;@1;)
      local.get 3
      local.get 1
      i32.store offset=24
      local.get 3
      local.get 1
      i32.store offset=20
      local.get 3
      local.get 1
      i32.store offset=16
      local.get 3
      local.get 1
      i32.store offset=12
      local.get 2
      i32.const -16
      i32.add
      local.get 1
      i32.store
      local.get 2
      i32.const -20
      i32.add
      local.get 1
      i32.store
      local.get 2
      i32.const -24
      i32.add
      local.get 1
      i32.store
      local.get 2
      i32.const -28
      i32.add
      local.get 1
      i32.store
      local.get 4
      local.get 3
      i32.const 4
      i32.and
      i32.const 24
      i32.or
      local.tee 5
      i32.sub
      local.tee 2
      i32.const 32
      i32.lt_u
      br_if 0 (;@1;)
      local.get 1
      i64.extend_i32_u
      i64.const 4294967297
      i64.mul
      local.set 6
      local.get 3
      local.get 5
      i32.add
      local.set 1
      loop  ;; label = @2
        local.get 1
        local.get 6
        i64.store offset=24
        local.get 1
        local.get 6
        i64.store offset=16
        local.get 1
        local.get 6
        i64.store offset=8
        local.get 1
        local.get 6
        i64.store
        local.get 1
        i32.const 32
        i32.add
        local.set 1
        local.get 2
        i32.const -32
        i32.add
        local.tee 2
        i32.const 31
        i32.gt_u
        br_if 0 (;@2;)
      end
    end
    local.get 0)
  (func (;22;) (type 1) (param i32 i32) (result i32)
    i32.const 0
    local.get 0
    i32.store offset=16778316
    i32.const 0
    local.get 1
    i32.store offset=16778312
    i32.const 0)
  (func (;23;) (type 3) (param i32 i32)
    (local i32 i32 i64 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    i32.const 1
    local.set 3
    block  ;; label = @1
      i32.const 0
      br_if 0 (;@1;)
      i32.const 0
      i64.load offset=16777232
      local.tee 4
      i32.wrap_i64
      i32.const 4
      i32.const 0
      i32.const 0
      local.get 4
      i64.const 32
      i64.shr_u
      i32.wrap_i64
      i32.load
      call_indirect (type 0)
      local.tee 5
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      local.get 5
      i32.store
      local.get 5
      local.get 1
      i32.load align=1
      i32.store align=1
      i32.const 0
      local.set 3
    end
    local.get 0
    local.get 3
    i32.store16 offset=4
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;24;) (type 7) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    local.get 0
    i32.const 16777255
    call 23
    local.get 0
    i32.load
    local.set 1
    local.get 0
    i32.const 8
    i32.add
    i32.const 16777251
    call 23
    local.get 1
    local.get 0
    i32.load offset=8
    call 3
    local.get 0
    i32.const 16
    i32.add
    global.set 0
    i32.const 16777240)
  (func (;25;) (type 0) (param i32 i32 i32 i32) (result i32)
    (local i32 i32 i32)
    i32.const 0
    local.set 4
    block  ;; label = @1
      i32.const -1
      local.get 1
      i32.const 4
      i32.add
      local.tee 5
      local.get 5
      local.get 1
      i32.lt_u
      select
      local.tee 1
      i32.const 1
      local.get 2
      i32.shl
      local.tee 2
      local.get 1
      local.get 2
      i32.gt_u
      select
      local.tee 2
      i32.const -1
      i32.add
      i32.clz
      local.tee 1
      i32.eqz
      br_if 0 (;@1;)
      block  ;; label = @2
        block  ;; label = @3
          i64.const 1
          i32.const 32
          local.get 1
          i32.sub
          i64.extend_i32_u
          i64.const 65535
          i64.and
          i64.shl
          i32.wrap_i64
          local.tee 5
          i32.ctz
          i32.const -3
          i32.add
          local.tee 1
          i32.const 13
          i32.ge_u
          br_if 0 (;@3;)
          local.get 1
          i32.const 2
          i32.shl
          local.tee 6
          i32.const 16778320
          i32.add
          local.tee 2
          i32.load
          local.tee 1
          i32.eqz
          br_if 1 (;@2;)
          local.get 2
          local.get 5
          local.get 1
          i32.add
          i32.const -4
          i32.add
          i32.load
          i32.store
          local.get 1
          return
        end
        local.get 2
        i32.const 65539
        i32.add
        i32.const 16
        i32.shr_u
        call 26
        local.set 4
        br 1 (;@1;)
      end
      block  ;; label = @2
        local.get 6
        i32.const 16778372
        i32.add
        local.tee 2
        i32.load
        local.tee 1
        i32.const 65535
        i32.and
        br_if 0 (;@2;)
        i32.const 1
        call 26
        local.tee 1
        i32.eqz
        br_if 1 (;@1;)
        local.get 2
        local.get 1
        local.get 5
        i32.add
        i32.store
        local.get 1
        return
      end
      local.get 2
      local.get 1
      local.get 5
      i32.add
      i32.store
      local.get 1
      return
    end
    local.get 4)
  (func (;26;) (type 6) (param i32) (result i32)
    (local i32 i32)
    block  ;; label = @1
      i64.const 1
      i32.const 32
      local.get 0
      i32.const -1
      i32.add
      i32.clz
      i32.sub
      i64.extend_i32_u
      i64.const 65535
      i64.and
      i64.shl
      i32.wrap_i64
      local.tee 1
      i32.ctz
      i32.const 2
      i32.shl
      i32.const 16778424
      i32.add
      local.tee 2
      i32.load
      local.tee 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 2
      local.get 1
      i32.const 16
      i32.shl
      local.get 0
      i32.add
      i32.const -4
      i32.add
      i32.load
      i32.store
      local.get 0
      return
    end
    i32.const 0
    local.get 1
    memory.grow
    local.tee 0
    i32.const 16
    i32.shl
    local.get 0
    i32.const -1
    i32.eq
    select)
  (func (;27;) (type 8) (param i32 i32 i32 i32 i32 i32) (result i32)
    (local i32)
    i32.const -1
    local.get 4
    i32.const 4
    i32.add
    local.tee 6
    local.get 6
    local.get 4
    i32.lt_u
    select
    local.tee 6
    i32.const 1
    local.get 3
    i32.shl
    local.tee 4
    local.get 6
    local.get 4
    i32.gt_u
    select
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        i64.const 1
        i32.const 32
        local.get 2
        i32.const 4
        i32.add
        local.tee 2
        local.get 4
        local.get 2
        local.get 4
        i32.gt_u
        select
        local.tee 4
        i32.const -1
        i32.add
        i32.clz
        i32.sub
        i64.extend_i32_u
        i64.const 65535
        i64.and
        i64.shl
        i32.wrap_i64
        local.tee 2
        i32.ctz
        i32.const -3
        i32.add
        i32.const 12
        i32.gt_u
        br_if 0 (;@2;)
        local.get 3
        i32.const -1
        i32.add
        i32.clz
        local.tee 4
        br_if 1 (;@1;)
        i32.const 0
        return
      end
      i64.const 1
      i32.const 32
      local.get 4
      i32.const 65539
      i32.add
      i32.const 16
      i32.shr_u
      i32.const -1
      i32.add
      i32.clz
      i32.sub
      i64.extend_i32_u
      i64.const 65535
      i64.and
      i64.shl
      i32.wrap_i64
      i64.const 1
      i32.const 32
      local.get 3
      i32.const 65539
      i32.add
      i32.const 16
      i32.shr_u
      i32.const -1
      i32.add
      i32.clz
      i32.sub
      i64.extend_i32_u
      i64.const 65535
      i64.and
      i64.shl
      i32.wrap_i64
      i32.eq
      return
    end
    local.get 2
    i64.const 1
    i32.const 32
    local.get 4
    i32.sub
    i64.extend_i32_u
    i64.const 65535
    i64.and
    i64.shl
    i32.wrap_i64
    i32.eq)
  (func (;28;) (type 9) (param i32 i32 i32 i32 i32)
    (local i32)
    block  ;; label = @1
      block  ;; label = @2
        i64.const 1
        i32.const 32
        local.get 2
        i32.const 4
        i32.add
        local.tee 2
        i32.const 1
        local.get 3
        i32.shl
        local.tee 3
        local.get 2
        local.get 3
        i32.gt_u
        select
        local.tee 3
        i32.const -1
        i32.add
        i32.clz
        i32.sub
        i64.extend_i32_u
        i64.const 65535
        i64.and
        i64.shl
        i32.wrap_i64
        local.tee 2
        i32.ctz
        i32.const -3
        i32.add
        local.tee 5
        i32.const 13
        i32.ge_u
        br_if 0 (;@2;)
        local.get 5
        i32.const 2
        i32.shl
        i32.const 16778320
        i32.add
        local.set 3
        local.get 1
        local.get 2
        i32.add
        i32.const -4
        i32.add
        local.set 2
        br 1 (;@1;)
      end
      i64.const 1
      i32.const 32
      local.get 3
      i32.const 65539
      i32.add
      i32.const 16
      i32.shr_u
      i32.const -1
      i32.add
      i32.clz
      i32.sub
      i64.extend_i32_u
      i64.const 65535
      i64.and
      i64.shl
      i32.wrap_i64
      local.tee 2
      i32.ctz
      i32.const 2
      i32.shl
      i32.const 16778424
      i32.add
      local.set 3
      local.get 1
      local.get 2
      i32.const 16
      i32.shl
      i32.add
      i32.const -4
      i32.add
      local.set 2
    end
    local.get 2
    local.get 3
    i32.load
    i32.store
    local.get 3
    local.get 1
    i32.store)
  (table (;0;) 4 4 funcref)
  (memory (;0;) 257)
  (global (;0;) (mut i32) (i32.const 16777216))
  (export "memory" (memory 0))
  (export "_start" (func 5))
  (export "malloc" (func 8))
  (export "Test" (func 24))
  (elem (;0;) (i32.const 1) func 25 27 28)
  (data (;0;) (i32.const 16777216) "\01\00\00\00\02\00\00\00\03\00\00\00\00\00\00\00\00\00\00\00\00\00\00\01HelloWorld\00514\00114\00")
  (data (;1;) (i32.const 16777264) "\00\00\00\000\00\00\010\00\00\01\00\00\00\00\00\00\00\00@\00\00\01@\00\00\01\00\00\00\00\00\00\00\00P\00\00\01P\00\00\01\00\00\00\00\00\00\00\00`\00\00\01`\00\00\01\00\00\00\00\00\00\00\00p\00\00\01p\00\00\01\00\00\00\00\00\00\00\00\80\00\00\01\80\00\00\01\00\00\00\00\00\00\00\00\90\00\00\01\90\00\00\01\00\00\00\00\00\00\00\00\a0\00\00\01\a0\00\00\01\00\00\00\00\00\00\00\00\b0\00\00\01\b0\00\00\01\00\00\00\00\00\00\00\00\c0\00\00\01\c0\00\00\01\00\00\00\00\00\00\00\00\d0\00\00\01\d0\00\00\01\00\00\00\00\00\00\00\00\e0\00\00\01\e0\00\00\01\00\00\00\00\00\00\00\00\f0\00\00\01\f0\00\00\01\00\00\00\00\00\00\00\00\00\01\00\01\00\01\00\01\00\00\00\00\00\00\00\00\10\01\00\01\10\01\00\01\00\00\00\00\00\00\00\00 \01\00\01 \01\00\01\00\00\00\00\00\00\00\000\01\00\010\01\00\01\00\00\00\00\00\00\00\00@\01\00\01@\01\00\01\00\00\00\00\00\00\00\00P\01\00\01P\01\00\01\00\00\00\00\00\00\00\00`\01\00\01`\01\00\01\00\00\00\00\00\00\00\00p\01\00\01p\01\00\01\00\00\00\00\00\00\00\00\80\01\00\01\80\01\00\01\00\00\00\00\00\00\00\00\90\01\00\01\90\01\00\01\00\00\00\00\00\00\00\00\a0\01\00\01\a0\01\00\01\00\00\00\00\00\00\00\00\b0\01\00\01\b0\01\00\01\00\00\00\00\00\00\00\00\c0\01\00\01\c0\01\00\01\00\00\00\00\00\00\00\00\d0\01\00\01\d0\01\00\01\00\00\00\00\00\00\00\00\e0\01\00\01\e0\01\00\01\00\00\00\00\00\00\00\00\f0\01\00\01\f0\01\00\01\00\00\00\00\00\00\00\00\00\02\00\01\00\02\00\01\00\00\00\00\00\00\00\00\10\02\00\01\10\02\00\01\00\00\00\00\00\00\00\00 \02\00\01 \02\00\01\00\00\00\00\00\00\00\000\02\00\010\02\00\01\00\00\00\00\00\00\00\00@\02\00\01@\02\00\01\00\00\00\00\00\00\00\00P\02\00\01P\02\00\01\00\00\00\00\00\00\00\00`\02\00\01`\02\00\01\00\00\00\00\00\00\00\00p\02\00\01p\02\00\01\00\00\00\00\00\00\00\00\80\02\00\01\80\02\00\01\00\00\00\00\00\00\00\00\90\02\00\01\90\02\00\01\00\00\00\00\00\00\00\00\a0\02\00\01\a0\02\00\01\00\00\00\00\00\00\00\00\b0\02\00\01\b0\02\00\01\00\00\00\00\00\00\00\00\c0\02\00\01\c0\02\00\01\00\00\00\00\00\00\00\00\d0\02\00\01\d0\02\00\01\00\00\00\00\00\00\00\00\e0\02\00\01\e0\02\00\01\00\00\00\00\00\00\00\00\f0\02\00\01\f0\02\00\01\00\00\00\00\00\00\00\00\00\03\00\01\00\03\00\01\00\00\00\00\00\00\00\00\10\03\00\01\10\03\00\01\00\00\00\00\00\00\00\00 \03\00\01 \03\00\01\00\00\00\00\00\00\00\000\03\00\010\03\00\01\00\00\00\00\00\00\00\00@\03\00\01@\03\00\01\00\00\00\00\00\00\00\00P\03\00\01P\03\00\01\00\00\00\00\00\00\00\00`\03\00\01`\03\00\01\00\00\00\00\00\00\00\00p\03\00\01p\03\00\01\00\00\00\00\00\00\00\00\80\03\00\01\80\03\00\01\00\00\00\00\00\00\00\00\90\03\00\01\90\03\00\01\00\00\00\00\00\00\00\00\a0\03\00\01\a0\03\00\01\00\00\00\00\00\00\00\00\b0\03\00\01\b0\03\00\01\00\00\00\00\00\00\00\00\c0\03\00\01\c0\03\00\01\00\00\00\00\00\00\00\00\d0\03\00\01\d0\03\00\01\00\00\00\00\00\00\00\00\e0\03\00\01\e0\03\00\01\00\00\00\00\00\00\00\00\f0\03\00\01\f0\03\00\01\00\00\00\00\00\00\00\00\00\04\00\01\00\04\00\01\00\00\00\00\00\00\00\00\10\04\00\01\10\04\00\01\00\00\00\00\00\00\00\00 \04\00\01 \04\00\01\00\00\00\00"))
